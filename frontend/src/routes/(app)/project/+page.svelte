<script lang="ts">
  import PageHeader from '$lib/components/PageHeader.svelte';

  let keyword = $state('');
  let page = $state(1);

  function search(v: string) {
    keyword = v;
    page = 1;
  }

  // ====== 静态资源（放在 frontend/static 下时，页面用 /xxx.svg 访问）======
  const ICON_ARROW = '/material-symbols--arrow-back-ios-new-rounded.svg';
  const ICON_CHECK = '/material-symbols--check-rounded.svg';
  const ICON_CLOSE = '/material-symbols--close.svg';
  const ICON_DATE = '/material-symbols--date-range-outline.svg';

  // ====== 顶部 7 个跳转按钮 ======
  type NavKey =
    | 'latest'
    | 'pigeon'
    | 'longest'
    | 'all'
    | 'translate'
    | 'proofread'
    | 'typeset';

  const navButtons: { key: NavKey; label: string; anchor: string }[] = [
    { key: 'latest', label: '最新项目', anchor: '#latest' },
    { key: 'pigeon', label: '鸽子救急', anchor: '#pigeon' },
    { key: 'longest', label: '咕了最久', anchor: '#longest' },
    { key: 'all', label: '全部项目', anchor: '#all' },
    { key: 'translate', label: '翻译专区', anchor: '#translate' },
    { key: 'proofread', label: '校对专区', anchor: '#proofread' },
    { key: 'typeset', label: '嵌字专区', anchor: '#typeset' }
  ];

  function jumpTo(anchor: string) {
    const el = document.querySelector(anchor);
    el?.scrollIntoView({ behavior: 'smooth', block: 'start' });
  }

  // ====== 公告栏（最多 10 张）======
  type Announcement = { id: string; imgUrl: string; alt?: string };
  let announcements = $state<Announcement[]>([
    { id: 'a1', imgUrl: 'https://picsum.photos/seed/ann1/520/280', alt: '公告 1' },
    { id: 'a2', imgUrl: 'https://picsum.photos/seed/ann2/520/280', alt: '公告 2' },
    { id: 'a3', imgUrl: 'https://picsum.photos/seed/ann3/520/280', alt: '公告 3' }
  ]);

  let annIndex = $state(0);
  function prevAnn() {
    if (!announcements.length) return;
    annIndex = (annIndex - 1 + announcements.length) % announcements.length;
  }
  function nextAnn() {
    if (!announcements.length) return;
    annIndex = (annIndex + 1) % announcements.length;
  }
  function goAnn(i: number) {
    annIndex = i;
  }

  // ====== 项目数据结构 ======
  type Project = {
    id: string;
    name: string;
    coverUrl: string;
    roles?: {
      translateFilled: boolean;
      proofreadFilled: boolean;
      typesetFilled: boolean;
    };
    lastUpdate?: string;
  };

  const demoProjects = (prefix: string, count = 6): Project[] =>
    Array.from({ length: count }).map((_, i) => ({
      id: `${prefix}-${i + 1}`,
      name: `桃园暗鬼 ${i + 1}`,
      coverUrl: `https://picsum.photos/seed/${prefix}-${i + 1}/260/260`,
      roles: {
        translateFilled: Math.random() > 0.5,
        proofreadFilled: Math.random() > 0.5,
        typesetFilled: Math.random() > 0.5
      }
    }));

  let latestProjects = $state<Project[]>(demoProjects('latest'));
  let pigeonProjects = $state<Project[]>(demoProjects('pigeon'));
  let translateZoneProjects = $state<Project[]>(demoProjects('tz'));
  let proofreadZoneProjects = $state<Project[]>(demoProjects('pz'));
  let typesetZoneProjects = $state<Project[]>(demoProjects('iz'));
  let allProjects = $state<Project[]>(demoProjects('all'));

  let longestProjects = $state<Project[]>(
    Array.from({ length: 6 }).map((_, i) => ({
      id: `long-${i + 1}`,
      name: `桃园暗鬼 ${i + 1}`,
      coverUrl: `https://picsum.photos/seed/long-${i + 1}/260/260`,
      lastUpdate: `2026.${(i % 12) + 1}.${(i % 27) + 1}`
    }))
  );

  function filterByKeyword(items: Project[]) {
    const k = keyword.trim();
    if (!k) return items;
    return items.filter((p) => p.name.includes(k) || p.id.includes(k));
  }

  function roleIcon(filled: boolean) {
    return {
      src: filled ? ICON_CHECK : ICON_CLOSE,
      className: filled ? 'role-icon ok' : 'role-icon no',
      alt: filled ? '已满' : '未满'
    };
  }
</script>

<PageHeader
  title="任务大厅"
  showSearch={true}
  searchValue={keyword}
  searchPlaceholder="查询项目ID 或 项目名"
  onSearch={search}
/>

<div class="page">
  <!-- 7 个跳转按钮 -->
  <div class="nav-buttons">
    {#each navButtons as b}
      <button class="nav-btn" type="button" on:click={() => jumpTo(b.anchor)}>
        {b.label}
      </button>
    {/each}
  </div>

  <!-- 第一排：4 个展示框 -->
  <section class="grid-4">
    <!-- 公告栏（高度固定，封面固定尺寸） -->
    <div class="panel panel-fixed-h" id="announce">
      <div class="announce-wrap">
        <button class="ann-arrow" type="button" on:click={prevAnn} aria-label="上一张公告">
          <img class="arrow" src={ICON_ARROW} alt="上一页" />
        </button>

        <div class="announce-card">
          {#if announcements.length > 0}
            <img
              class="announce-img"
              src={announcements[annIndex].imgUrl}
              alt={announcements[annIndex].alt ?? `公告 ${annIndex + 1}`}
            />
          {:else}
            <div class="announce-empty">暂无公告</div>
          {/if}

          <div class="ann-dots">
            {#each Array(10) as _, i}
              {#if i < announcements.length}
                <button
                  type="button"
                  class="dot {i === annIndex ? 'active' : ''}"
                  on:click={() => goAnn(i)}
                  aria-label={`切换到公告 ${i + 1}`}
                />
              {:else}
                <span class="dot placeholder" aria-hidden="true" />
              {/if}
            {/each}
          </div>
        </div>

        <button class="ann-arrow" type="button" on:click={nextAnn} aria-label="下一张公告">
          <img class="arrow rotate" src={ICON_ARROW} alt="下一页" />
        </button>
      </div>
    </div>

    <!-- 最新项目（高度固定） -->
    <div class="panel panel-fixed-h" id="latest">
      <div class="panel-head">
        <div class="panel-title">最新项目</div>
      </div>
      <div class="divider" />
      <div class="cards-6">
        {#each filterByKeyword(latestProjects).slice(0, 6) as p (p.id)}
          <div class="project-card">
            <div class="pc-left">
              <img class="pc-cover" src={p.coverUrl} alt={p.name} />
              <div class="pc-name" title={p.name}>{p.name}</div>
            </div>

            <div class="pc-right">
              <div class="role-row">
                <span class="role-label">翻译</span>
                <img
                  {...roleIcon(p.roles?.translateFilled ?? false)}
                  class={roleIcon(p.roles?.translateFilled ?? false).className}
                />
              </div>
              <div class="role-row">
                <span class="role-label">校对</span>
                <img
                  {...roleIcon(p.roles?.proofreadFilled ?? false)}
                  class={roleIcon(p.roles?.proofreadFilled ?? false).className}
                />
              </div>
              <div class="role-row">
                <span class="role-label">嵌字</span>
                <img
                  {...roleIcon(p.roles?.typesetFilled ?? false)}
                  class={roleIcon(p.roles?.typesetFilled ?? false).className}
                />
              </div>
            </div>
          </div>
        {/each}
      </div>
    </div>

    <!-- 鸽子救急（高度固定） -->
    <div class="panel panel-fixed-h" id="pigeon">
      <div class="panel-head">
        <div class="panel-title">鸽子救急</div>
      </div>
      <div class="divider" />
      <div class="cards-6">
        {#each filterByKeyword(pigeonProjects).slice(0, 6) as p (p.id)}
          <div class="project-card">
            <div class="pc-left">
              <img class="pc-cover" src={p.coverUrl} alt={p.name} />
              <div class="pc-name" title={p.name}>{p.name}</div>
            </div>

            <div class="pc-right">
              <div class="role-row">
                <span class="role-label">翻译</span>
                <img
                  {...roleIcon(p.roles?.translateFilled ?? false)}
                  class={roleIcon(p.roles?.translateFilled ?? false).className}
                />
              </div>
              <div class="role-row">
                <span class="role-label">校对</span>
                <img
                  {...roleIcon(p.roles?.proofreadFilled ?? false)}
                  class={roleIcon(p.roles?.proofreadFilled ?? false).className}
                />
              </div>
              <div class="role-row">
                <span class="role-label">嵌字</span>
                <img
                  {...roleIcon(p.roles?.typesetFilled ?? false)}
                  class={roleIcon(p.roles?.typesetFilled ?? false).className}
                />
              </div>
            </div>
          </div>
        {/each}
      </div>
    </div>

    <!-- 咕了最久（高度固定） -->
    <div class="panel panel-fixed-h" id="longest">
      <div class="panel-head">
        <div class="panel-title">咕了最久</div>
      </div>
      <div class="divider" />
      <div class="cards-6">
        {#each filterByKeyword(longestProjects).slice(0, 6) as p (p.id)}
          <div class="project-card longest">
            <div class="pc-left">
              <img class="pc-cover" src={p.coverUrl} alt={p.name} />
              <div class="pc-name" title={p.name}>{p.name}</div>
            </div>

            <div class="pc-right longest-right">
              <div class="last-title">上次更新</div>
              <div class="last-row">
                <img class="date-icon" src={ICON_DATE} alt="日期" />
                <span class="last-date">{p.lastUpdate ?? '-'}</span>
              </div>
            </div>
          </div>
        {/each}
      </div>
    </div>
  </section>

  <!-- 第二排：3 个展示框（高度也固定一致） -->
  <section class="grid-3">
    <div class="panel panel-fixed-h" id="translate">
      <div class="panel-head">
        <div class="panel-title">翻译专区</div>
      </div>
      <div class="divider" />
      <div class="cards-6">
        {#each filterByKeyword(translateZoneProjects).slice(0, 6) as p (p.id)}
          <div class="project-card">
            <div class="pc-left">
              <img class="pc-cover" src={p.coverUrl} alt={p.name} />
              <div class="pc-name" title={p.name}>{p.name}</div>
            </div>

            <div class="pc-right">
              <div class="role-row">
                <span class="role-label">翻译</span>
                <img
                  {...roleIcon(p.roles?.translateFilled ?? false)}
                  class={roleIcon(p.roles?.translateFilled ?? false).className}
                />
              </div>
              <div class="role-row">
                <span class="role-label">校对</span>
                <img
                  {...roleIcon(p.roles?.proofreadFilled ?? false)}
                  class={roleIcon(p.roles?.proofreadFilled ?? false).className}
                />
              </div>
              <div class="role-row">
                <span class="role-label">嵌字</span>
                <img
                  {...roleIcon(p.roles?.typesetFilled ?? false)}
                  class={roleIcon(p.roles?.typesetFilled ?? false).className}
                />
              </div>
            </div>
          </div>
        {/each}
      </div>
    </div>

    <div class="panel panel-fixed-h" id="proofread">
      <div class="panel-head">
        <div class="panel-title">校对专区</div>
      </div>
      <div class="divider" />
      <div class="cards-6">
        {#each filterByKeyword(proofreadZoneProjects).slice(0, 6) as p (p.id)}
          <div class="project-card">
            <div class="pc-left">
              <img class="pc-cover" src={p.coverUrl} alt={p.name} />
              <div class="pc-name" title={p.name}>{p.name}</div>
            </div>

            <div class="pc-right">
              <div class="role-row">
                <span class="role-label">翻译</span>
                <img
                  {...roleIcon(p.roles?.translateFilled ?? false)}
                  class={roleIcon(p.roles?.translateFilled ?? false).className}
                />
              </div>
              <div class="role-row">
                <span class="role-label">校对</span>
                <img
                  {...roleIcon(p.roles?.proofreadFilled ?? false)}
                  class={roleIcon(p.roles?.proofreadFilled ?? false).className}
                />
              </div>
              <div class="role-row">
                <span class="role-label">嵌字</span>
                <img
                  {...roleIcon(p.roles?.typesetFilled ?? false)}
                  class={roleIcon(p.roles?.typesetFilled ?? false).className}
                />
              </div>
            </div>
          </div>
        {/each}
      </div>
    </div>

    <div class="panel panel-fixed-h" id="typeset">
      <div class="panel-head">
        <div class="panel-title">嵌字专区</div>
      </div>
      <div class="divider" />
      <div class="cards-6">
        {#each filterByKeyword(typesetZoneProjects).slice(0, 6) as p (p.id)}
          <div class="project-card">
            <div class="pc-left">
              <img class="pc-cover" src={p.coverUrl} alt={p.name} />
              <div class="pc-name" title={p.name}>{p.name}</div>
            </div>

            <div class="pc-right">
              <div class="role-row">
                <span class="role-label">翻译</span>
                <img
                  {...roleIcon(p.roles?.translateFilled ?? false)}
                  class={roleIcon(p.roles?.translateFilled ?? false).className}
                />
              </div>
              <div class="role-row">
                <span class="role-label">校对</span>
                <img
                  {...roleIcon(p.roles?.proofreadFilled ?? false)}
                  class={roleIcon(p.roles?.proofreadFilled ?? false).className}
                />
              </div>
              <div class="role-row">
                <span class="role-label">嵌字</span>
                <img
                  {...roleIcon(p.roles?.typesetFilled ?? false)}
                  class={roleIcon(p.roles?.typesetFilled ?? false).className}
                />
              </div>
            </div>
          </div>
        {/each}
      </div>
    </div>
  </section>

  <!-- 第三排：全部项目（也固定高度） -->
  <section class="grid-1">
    <div class="panel panel-fixed-h" id="all">
      <div class="panel-head">
        <div class="panel-title">全部项目</div>
      </div>
      <div class="divider" />
      <div class="cards-6 wide">
        {#each filterByKeyword(allProjects).slice(0, 6) as p (p.id)}
          <div class="project-card">
            <div class="pc-left">
              <img class="pc-cover" src={p.coverUrl} alt={p.name} />
              <div class="pc-name" title={p.name}>{p.name}</div>
            </div>

            <div class="pc-right">
              <div class="role-row">
                <span class="role-label">翻译</span>
                <img
                  {...roleIcon(p.roles?.translateFilled ?? false)}
                  class={roleIcon(p.roles?.translateFilled ?? false).className}
                />
              </div>
              <div class="role-row">
                <span class="role-label">校对</span>
                <img
                  {...roleIcon(p.roles?.proofreadFilled ?? false)}
                  class={roleIcon(p.roles?.proofreadFilled ?? false).className}
                />
              </div>
              <div class="role-row">
                <span class="role-label">嵌字</span>
                <img
                  {...roleIcon(p.roles?.typesetFilled ?? false)}
                  class={roleIcon(p.roles?.typesetFilled ?? false).className}
                />
              </div>
            </div>
          </div>
        {/each}
      </div>
    </div>
  </section>
</div>

<style>
  /* ===== 可调参数：统一封面尺寸 & 面板统一高度 ===== */
  :global(:root) {
    --panel-fixed-height: 520px; /* ✅ 公告栏 & 各专区/展示框统一高度（不随内容变化） */
    --announce-cover-w: 100%;
    --announce-cover-h: 320px; /* ✅ 公告封面固定高度 */

    --project-cover-size: 92px; /* ✅ 每个项目封面固定尺寸（所有区一致） */
  }

  .page {
    padding: 18px 18px 40px;
  }

  /* ===== 顶部按钮 ===== */
  .nav-buttons {
    display: flex;
    justify-content: center;
    gap: 14px;
    flex-wrap: wrap;
    margin: 6px 0 18px;
  }

  .nav-btn {
    width: clamp(88px, 25vw, 252px);
    background: #ffffff;
    border: 1px solid #cfcfc9;
    border-radius: 8px;
    padding: 10px 16px;
    font-size: 14px;
    cursor: pointer;
    line-height: 1;
    user-select: none;
  }

  .nav-btn:hover {
    filter: brightness(0.98);
  }

  /* ===== 面板容器（固定高度）===== */
  .panel {
    background: #ffffff;
    border: 1px solid #e2e2dd;
    border-radius: 10px;
    padding: 12px;
    min-width: 0;
  }

  .panel-fixed-h {
    height: var(--panel-fixed-height);      /* ✅ 统一高度 */
    box-sizing: border-box;
    overflow: hidden;                      /* ✅ 内容不够不影响高度；太多则裁剪 */
    display: flex;
    flex-direction: column;
  }

  .panel-head {
    display: flex;
    align-items: center;
    justify-content: space-between;
    gap: 10px;
    padding: 2px 4px 8px;
    flex: 0 0 auto;
  }

  .panel-title {
    font-size: 14px;
    font-weight: 600;
    color: #111;
  }

  .divider {
    height: 1px;
    background: #cfcfc9;
    margin: 0 2px 10px;
    flex: 0 0 auto;
  }

  /* 面板内容区：占满剩余高度（确保内部排版稳定） */
  .panel-fixed-h :global(.cards-6),
  .panel-fixed-h :global(.cards-6.wide),
  .panel-fixed-h .announce-wrap {
    flex: 1 1 auto;
    min-height: 0;
    overflow: auto; /* ✅ 内容超过固定高度时可滚动（不影响外层高度一致） */
  }

  /* ===== 栅格布局 ===== */
  .grid-4 {
    display: grid;
    grid-template-columns: repeat(4, minmax(0, 1fr));
    gap: 14px;
    align-items: start;
    margin-bottom: 14px;
  }

  .grid-3 {
    display: grid;
    grid-template-columns: repeat(3, minmax(0, 1fr));
    gap: 14px;
    align-items: start;
    margin-bottom: 14px;
  }

  .grid-1 {
    display: grid;
    grid-template-columns: 1fr;
    gap: 14px;
    align-items: start;
  }

  /* ===== 公告栏（封面固定尺寸）===== */
  .announce-wrap {
    display: grid;
    grid-template-columns: 34px 1fr 34px;
    align-items: center;
    gap: 10px;
  }

  .ann-arrow {
    width: 34px;
    height: 34px;
    border: 1px solid #cfcfc9;
    background: #ffffff;
    border-radius: 8px;
    display: grid;
    place-items: center;
    cursor: pointer;
    flex: 0 0 auto;
  }

  .ann-arrow:hover {
    filter: brightness(0.98);
  }

  .arrow {
    width: 18px;
    height: 18px;
    display: block;
  }
  .arrow.rotate {
    transform: rotate(180deg);
  }

  .announce-card {
    position: relative;
    border: 1px solid #cfcfc9;
    border-radius: 10px;
    overflow: hidden;
    background: #ffffff;

    /* ✅ 固定公告封面区域高度 */
    width: var(--announce-cover-w);
    height: var(--announce-cover-h);
    display: grid;
    place-items: center;
  }

  .announce-img {
    width: 100%;
    height: 100%;
    object-fit: cover; /* ✅ 固定尺寸下裁剪填充 */
    display: block;
  }

  .announce-empty {
    color: #666;
    font-size: 14px;
    padding: 28px;
  }

  .ann-dots {
    position: absolute;
    left: 50%;
    bottom: 10px;
    transform: translateX(-50%);
    display: flex;
    gap: 6px;
    align-items: center;
  }

  .dot {
    width: 8px;
    height: 8px;
    border-radius: 999px;
    background: #ffffff;
    border: 1px solid #cfcfc9;
    cursor: pointer;
    padding: 0;
  }

  .dot.active {
    transform: scale(1.15);
  }

  .dot.placeholder {
    opacity: 0.25;
    cursor: default;
  }

  /* ===== 项目卡片（统一封面尺寸）===== */
  .cards-6 {
    display: grid;
    grid-template-columns: repeat(2, minmax(0, 1fr));
    gap: 10px;
  }

  .cards-6.wide {
    grid-template-columns: repeat(3, minmax(0, 1fr));
  }

  .project-card {
    background: #ffffff;
    border: 1px solid #e2e2dd;
    border-radius: 10px;
    padding: 10px;
    display: grid;

    /* ✅ 左侧封面固定宽度，右侧信息固定宽度 */
    grid-template-columns: var(--project-cover-size) 1fr;
    gap: 10px;
    align-items: center;
    min-width: 0;
  }

  .pc-left {
    width: var(--project-cover-size);
    display: grid;
    grid-template-rows: var(--project-cover-size) auto;
    gap: 8px;
    justify-items: center;
  }

  .pc-cover {
    width: var(--project-cover-size);
    height: var(--project-cover-size); /* ✅ 强制统一 */
    border-radius: 8px;
    object-fit: cover;
    border: 1px solid #e2e2dd;
    background: #fff;
    display: block;
  }

  .pc-name {
    width: 100%;
    font-size: 13px;
    color: #111;
    line-height: 1.2;
    text-align: center;
    white-space: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;
  }

  /* 右侧信息（岗位列 / 上次更新） */
  .pc-right {
    display: grid;
    gap: 8px;
    padding-top: 2px;
    align-content: start;
    min-width: 0;
  }

  .role-row {
    display: flex;
    align-items: center;
    justify-content: space-between;
    gap: 8px;
  }

  .role-label {
    font-size: 12px;
    color: #333;
  }

  .role-icon {
    width: 18px;
    height: 18px;
    display: block;
  }

  /* ✅ check 绿色 / close 红色 */
  .role-icon.ok {
    filter: hue-rotate(90deg) saturate(4) brightness(0.9);
  }
  .role-icon.no {
    filter: hue-rotate(330deg) saturate(6) brightness(0.9);
  }

  /* ===== 咕了最久：右侧显示上次更新 ===== */
  .project-card.longest {
    grid-template-columns: var(--project-cover-size) 1fr;
  }

  .longest-right {
    display: grid;
    gap: 6px;
    align-content: start;
  }

  .last-title {
    font-size: 12px;
    color: #555;
  }

  .last-row {
    display: flex;
    align-items: center;
    gap: 8px;
  }

  .date-icon {
    width: 18px;
    height: 18px;
    display: block;
  }

  .last-date {
    font-size: 12px;
    color: #111;
    white-space: nowrap;
  }

  /* ===== 响应式 ===== */
  @media (max-width: 1280px) {
    .grid-4 {
      grid-template-columns: repeat(2, minmax(0, 1fr));
    }
  }

  @media (max-width: 860px) {
    .grid-3 {
      grid-template-columns: 1fr;
    }
    .cards-6 {
      grid-template-columns: 1fr;
    }
    .cards-6.wide {
      grid-template-columns: 1fr;
    }
  }
</style>
