<script lang="ts">
  import PageHeader from '$lib/components/PageHeader.svelte';

  // 你原本引入的选择器我先保留（后面接真实逻辑时用得上）
  import UserPicker, { type User } from '$lib/components/Picker/UserPicker.svelte';
  import ChapterPicker, { type Chapter } from '$lib/components/Picker/ChapterPicker.svelte';

  // ====== 示例数据（你接接口时替换即可）======
  type ProjectInfo = {
    translated_name: string;
    name: string;
    author: string;
    status: number; // 1 连载中
    updated_time: string;
    tags: string;
    translation_description: string;
    description: string;
    cover_url: string;
  };

  const project = $state<ProjectInfo>({
    translated_name: '间谍过家家第三季',
    name: 'SPY×FAMILY',
    author: '远藤达哉',
    status: 1,
    updated_time: '2026-1-9 千棠 翻译 121话',
    tags: '治愈',
    translation_description:
      '翻译简介：机翻简介：结城真司是一名普通的上班族：35岁，单身，没有女朋友，还和父母住在一起。他为独自生活的未来感到焦虑，但恋爱经验匮乏，害怕结婚。然而，一次意外让他邂逅了医生番绘里子，并坠入了无法想象的爱情，彻底改变了他的生活！人气成人恋爱漫画《公司里有我喜欢的人》的作者榎本赤丸，为您带来30多岁男士的恋爱训练！！',
    description:
      '生肉简介：35歳独身、彼女なし、実家暮らし」の平凡なサラリーマン・悠木慎史。独りで生きる未来に不安があるけれど、恋愛経験は少ないし婚活も怖い――。そんな悠木が、とあるアクシデントで医師・伴瑛理子と出会い、人生がひっくり返る「想像以上の恋」に落ちる！ 　大ヒット社会人恋愛漫画『この会社に好きな人がいます』の榎本あかまるが贈る、ミドサーから始める恋愛トレーニング!!',
    cover_url: 'https://picsum.photos/seed/cover/420/560'
  });

  function statusText(s: number) {
    if (s === 1) return '连载中';
    if (s === 2) return '已完结';
    return '未知';
  }

  // ====== 人员信息 ======
  type Person = { id: string; name: string; avatarUrl?: string };

  const makePeople = (seed: string, n: number): Person[] =>
    Array.from({ length: n }).map((_, i) => ({
      id: `${seed}-${i + 1}`,
      name: `工鸣蛇${i + 1}`,
      avatarUrl: `https://picsum.photos/seed/${seed}-${i + 1}/80/80`
    }));

  type RoleGroup = { title: string; people: Person[] };

  const roleGroups = $state<RoleGroup[]>([
    { title: '管理员', people: makePeople('admin', 1) },
    { title: '图源', people: makePeople('raw', 1) },
    { title: '翻译', people: makePeople('tl', 1) },
    { title: '校对', people: makePeople('pr', 2) },
    { title: '嵌字', people: makePeople('ts', 2) },
    { title: '监督', people: makePeople('su', 2) }
  ]);

  // ====== 展示章节 & 图片 ======
  type ChapterBlock = { id: string; title: string; images: string[] };

  const makeImages = (seed: string, count: number) =>
    Array.from({ length: count }).map((_, i) => `https://picsum.photos/seed/${seed}-${i + 1}/240/320`);

  const chapters = $state<ChapterBlock[]>([
    { id: 'c1', title: '1.1', images: makeImages('c11', 15) },
    { id: 'c2', title: '1.2', images: makeImages('c12', 15) }
  ]);
</script>

<PageHeader title="修改项目">
  <div slot="actions" class="page-actions">
    <button class="primary">参与项目</button>
    <button class="ghost">订阅</button>
  </div>
</PageHeader>

<div class="page">
  <!-- 顶部：两栏 6:4 -->
  <section class="top-grid">
    <!-- 左：项目信息 -->
    <div class="box">
      <div class="box-title">项目信息</div>
      <div class="box-divider" />

      <div class="project-main">
        <div class="cover-wrap">
          <img class="cover" src={project.cover_url} alt="项目封面" />
        </div>

        <div class="info-wrap">
          <div class="p-translated">{project.translated_name}</div>
          <div class="p-name">{project.name}</div>

          <div class="row-2">
            <div class="kv">
              <span class="k">作者：</span><span class="v">{project.author}</span>
            </div>
            <div class="kv">
              <span class="k">状态：</span><span class="v">{statusText(project.status)}</span>
            </div>
          </div>

          <div class="kv line">
            <span class="k">更新时间：</span><span class="v">{project.updated_time}</span>
          </div>

          <div class="kv line">
            <span class="k">类别：</span><span class="v">{project.tags}</span>
          </div>

          <div class="block">
            <div class="block-text">{project.translation_description}</div>
          </div>

          <div class="block">
            <div class="block-text">{project.description}</div>
          </div>
        </div>
      </div>
    </div>

    <!-- 右：人员信息 -->
    <div class="box">
      <div class="box-title">人员信息</div>
      <div class="box-divider" />

      <div class="roles">
        {#each roleGroups as g (g.title)}
          <div class="role-section">
            <div class="role-head">
              <div class="role-title">{g.title}</div>
              <div class="count-pill">{g.people.length}</div>
            </div>

            <div class="people-row">
              {#each g.people as u (u.id)}
                <div class="person">
                  <div class="avatar">
                    {#if u.avatarUrl}
                      <img src={u.avatarUrl} alt={u.name} />
                    {:else}
                      <span class="avatar-fallback">工</span>
                    {/if}
                  </div>
                  <div class="person-name">{u.name}</div>
                </div>
              {/each}
            </div>

            <div class="thin-divider" />
          </div>
        {/each}
      </div>
    </div>
  </section>

  <!-- 底部：展示章节 -->
  <section class="bottom">
    <div class="box">
      <div class="box-title">展示章节</div>
      <div class="box-divider" />

      <div class="chapter-list">
        {#each chapters as ch (ch.id)}
          <div class="chapter-card">
            <div class="chapter-title">{ch.title}</div>
            <div class="chapter-divider" />

            <div class="image-strip">
              {#each ch.images as img (img)}
                <div class="img-cell">
                  <img class="manga-img" src={img} alt="章节图片" />
                </div>
              {/each}
            </div>
          </div>
        {/each}
      </div>
    </div>
  </section>
</div>

<style>
  /* ===== 页面整体 ===== */
  .page {
    padding: 14px 18px 40px;
  }

  .page-actions {
    display: flex;
    gap: 10px;
    align-items: center;
  }

  button.primary {
    background: #1f6feb;
    color: #fff;
    border: 1px solid #1f6feb;
    border-radius: 8px;
    padding: 8px 14px;
    cursor: pointer;
  }

  button.ghost {
    background: #ffffff;
    color: #111;
    border: 1px solid #cfcfc9;
    border-radius: 8px;
    padding: 8px 14px;
    cursor: pointer;
  }

  /* ===== 两栏 6:4 ===== */
  .top-grid {
    display: grid;
    grid-template-columns: 6fr 4fr;
    gap: 14px;
    align-items: start;
    margin-bottom: 14px;
    align-items: stretch;
  }

  /* ===== 通用大框：白底 + 797979 边框 ===== */
  .box {
    background: #ffffff;
    border: 1px solid #E2E2DD;
    border-radius: 6px;
    padding: 12px;
    box-sizing: border-box;
  }

  .box-title {
    font-size: 14px;
    font-weight: 600;
    color: #5f5f5f;
    padding: 2px 4px 8px;
  }

  .box-divider {
    height: 1px;
    background: #cfcfc9;
    margin: 0 2px 12px;
  }

  /* ===== 左侧：项目信息内容 ===== */
  .project-main {
    margin: 25px auto;  
    max-width: 600px; 
    display: grid;
    grid-template-columns: 220px 1fr;
    gap: 18px;
    align-items: start;
  }

  /* 封面固定尺寸（你要一致） */
  .cover-wrap {
    width: 220px;
  }
  .cover {
    width: 220px;
    height: 300px;
    object-fit: cover;
    display: block;
    border-radius: 4px;
    border: 1px solid #e2e2dd;
    background: #fff;
  }

  .info-wrap {
    color: #5f5f5f;
    font-size: 13px;
    line-height: 1.55;
    min-width: 0;
  }

  .p-translated {
    font-size: 16px;
    font-weight: 600;
    color: #5f5f5f;
    margin-bottom: 6px;
  }

  .p-name {
    font-size: 13px;
    color: #5f5f5f;
    margin-bottom: 10px;
  }

  .row-2 {
    display: grid;
    grid-template-columns: 1fr 1fr;
    gap: 10px;
    margin-bottom: 8px;
  }

  .kv .k {
    color: #5f5f5f;
    margin-right: 4px;
    white-space: nowrap;
  }
  .kv .v {
    color: #5f5f5f;
  }

  .kv.line {
    margin-bottom: 8px;
  }

  .block {
    margin-top: 10px;
  }

  .block-text {
    white-space: pre-wrap;
    word-break: break-word;
    color: #5f5f5f;
  }

  /* ===== 右侧：人员信息 ===== */
  .roles {
    display: flex;
    flex-direction: column;
    gap: 10px;
  }

  .role-section {
    display: flex;
    flex-direction: column;
    gap: 10px;
  }

  .role-head {
    display: flex;
    align-items: center;
    gap: 8px;
    padding: 0 2px;
  }

  .role-title {
    font-size: 13px;
    font-weight: 600;
    color: #5f5f5f;
  }

  .count-pill {
    width: 22px;
    height: 22px;
    border-radius: 999px;
    background: #ffffff;
    border: 1px solid #e2e2dd;
    color: #5f5f5f;
    display: grid;
    place-items: center;
    font-size: 12px;
    line-height: 1;
  }

  .people-row {
    display: flex;
    gap: 10px;
    flex-wrap: wrap;
    padding: 0 2px;
  }

  .person {
    display: flex;
    align-items: center;
    gap: 8px;
    padding: 6px 8px;
    border-radius: 6px;
  }

  .avatar {
    width: 32px;
    height: 32px;
    border-radius: 999px;
    border: 1px solid #e2e2dd;
    background: #fff;
    display: grid;
    place-items: center;
    overflow: hidden;
  }

  .avatar img {
    width: 100%;
    height: 100%;
    object-fit: cover;
    display: block;
  }

  .avatar-fallback {
    color: #5f5f5f;
    font-size: 14px;
  }

  .person-name {
    color: #5f5f5f;
    font-size: 12px;
    white-space: nowrap;
  }

  .thin-divider {
    height: 1px;
    background: #cfcfc9;
    margin: 0 2px;
  }

  /* ===== 底部：展示章节 ===== */
  .bottom .box {
    /* 跟上面同样白底/797979 */
  }

  .chapter-list {
    display: flex;
    flex-direction: column;
    gap: 12px;
  }

  .chapter-card {
    background: #ffffff;
    border: 1px solid #e2e2dd;
    border-radius: 6px;
    padding: 10px;
  }

  .chapter-title {
    font-size: 12px;
    color: #5f5f5f;
    font-weight: 600;
    padding: 2px 2px 8px;
  }

  .chapter-divider {
    height: 1px;
    background: #e2e2dd;
    margin: 0 2px 10px;
  }

  /* 图片横向排列，固定尺寸一致（15张左右可横向滚动） */
  .image-strip {
    display: flex;
    gap: 10px;
    overflow-x: auto;
    padding: 4px 2px 6px;
  }

  .img-cell {
    flex: 0 0 auto;
    width: 72px;   /* ✅ 每张图片固定大小（你要一致） */
    height: 92px;  /* ✅ 每张图片固定大小（你要一致） */
    border: 1px solid #e2e2dd;
    border-radius: 4px;
    background: #fff;
    overflow: hidden;
  }

  .manga-img {
    width: 100%;
    height: 100%;
    object-fit: cover;
    display: block;
  }

  /* ===== 响应式：窄屏改为上下 ===== */
  @media (max-width: 1080px) {
    .top-grid {
      grid-template-columns: 1fr;
    }
    .project-main {
      grid-template-columns: 180px 1fr;
    }
    .cover-wrap,
    .cover {
      width: 180px;
    }
    .cover {
      height: 250px;
    }
  }
</style>
