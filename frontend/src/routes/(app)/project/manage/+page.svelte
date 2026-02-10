<script lang="ts">
  import PageHeader from '$lib/components/PageHeader.svelte';

  // ===== 静态资源（放在 frontend/static 下时，页面用 /xxx.svg 访问）======
  const ICON_GROUPS = '/material-symbols--groups-outline.svg';
  const ICON_PERSON = '/material-symbols--person-outline.svg';

  type Project = {
    id: string;
    coverUrl: string;
    translated_name: string;
    name: string;
    published: boolean;
    description: string;
  };

  // 10 个模板项目
  let projects: Project[] = Array.from({ length: 10 }, (_, i) => ({
    id: `p${i + 1}`,
    coverUrl: `https://picsum.photos/seed/proj-mgr-${i + 1}/220/300`,
    translated_name: '间谍过家家第三季',
    name: 'SPY×FAMILY',
    published: i % 2 === 1,
    description:
      '翻译简介：机翻简介：结城真司是一名普通的上班族：35岁，单身，没有女朋友，还和父母住在一起。他为独自生活的未来感到焦虑，但恋爱经验匮乏，害怕结婚。然而，一次意外让他邂逅了医生番绘里子，并坠入了无法想象的爱情，彻底改变了他的生活！人气成人恋爱漫画《公司里有我喜欢的人》的作者榎本赤丸，为您带来30多岁男士的恋爱训练！！。'
  }));

  type NavKey = 'all' | 'joined';
  let activeNav: NavKey = 'all';

  $: pageTitle = activeNav === 'all' ? '组内全部项目' : '我加入的项目';
  $: headerCountText = `${projects.length} 个项目`;
</script>

<PageHeader title="项目管理">
  <div slot="actions" class="page-actions"></div>
</PageHeader>

<div class="page">
  <!-- 左：导航栏 -->
  <aside class="sidebar">
    <button
      type="button"
      class="nav-item"
      class:active={activeNav === 'all'}
      on:click={() => (activeNav = 'all')}
    >
      <img class="nav-icon" src={ICON_GROUPS} alt="" />
      <span class="nav-text">组内全部项目</span>
    </button>

    <button
      type="button"
      class="nav-item"
      class:active={activeNav === 'joined'}
      on:click={() => (activeNav = 'joined')}
    >
      <img class="nav-icon" src={ICON_PERSON} alt="" />
      <span class="nav-text">我加入的项目</span>
    </button>
  </aside>

  <!-- 中间分隔线 -->
  <div class="splitter" />

  <!-- 右：项目列表 -->
  <main class="content">
    <div class="content-title">{pageTitle}</div>

    <div class="list-card">
      <!-- 表头 -->
      <div class="list-head">
        <div class="list-count">{headerCountText}</div>
      </div>

      <!-- 项目列表 -->
      <div class="list-body">
        {#each projects as p (p.id)}
          <div class="proj-row">
            <img class="cover" src={p.coverUrl} alt="项目封面" />

            <div class="info">
              <div class="line1">
                <div class="tname">{p.translated_name}</div>
                <div class="badge">{p.published ? '已发布' : '未发布'}</div>
              </div>

              <div class="oname">{p.name}</div>

              <div class="desc">{p.description}</div>
            </div>
          </div>
        {/each}
      </div>
    </div>
  </main>
</div>

<style>
  .page {
    display: grid;
    grid-template-columns: 240px 1px 1fr;
    gap: 0;
    padding: 16px 20px;
    box-sizing: border-box;
    min-height: calc(100vh - 72px);
    background: #f6f5f1;
  }

  /* 左侧导航 */
  .sidebar {
    display: flex;
    flex-direction: column;
    gap: 12px;
    padding: 10px 10px 10px 0;
    box-sizing: border-box;
  }

  .nav-item {
    height: 40px;
    border: none;
    background: transparent;
    display: flex;
    align-items: center;
    gap: 10px;
    cursor: pointer;
    padding: 0 8px;
    border-radius: 10px;
  }

  .nav-item:hover {
    background: rgba(0, 0, 0, 0.03);
  }

  .nav-item.active {
    background: rgba(0, 0, 0, 0.05);
  }

  .nav-icon {
    width: 18px;
    height: 18px;
    display: block;
    opacity: 0.9;
  }

  .nav-text {
    color: #2e2e2e;
    font-size: 13px;
  }

  /* 中间分隔线 */
  .splitter {
    width: 1px;
    background: #cfcfc9;
  }

  /* 右侧内容区 */
  .content {
    padding: 6px 0 0 18px;
    box-sizing: border-box;
    min-width: 0;
    display: flex;
    flex-direction: column;
    gap: 10px;
  }

  .content-title {
    color: #2e2e2e;
    font-size: 13px;
    font-weight: 700;
  }

  /* 列表卡片 */
  .list-card {
    border: 1px solid #e2e2dd;
    background: #ffffff;
    border-radius: 6px;
    overflow: hidden;
    min-width: 0;
  }

  /* 表头 */
  .list-head {
    height: 60px;
    background: #fbfbf9;
    border-bottom: 1px solid #e2e2dd;
    display: flex;
    align-items: center;
    padding: 0 16px;
    box-sizing: border-box;
  }

  .list-count {
    color: #2e2e2e;
    font-size: 13px;
    font-weight: 700;
  }

  /* 列表主体 */
  .list-body {
    display: flex;
    flex-direction: column;
  }

  .proj-row {
    display: grid;
    grid-template-columns: 86px 1fr;
    gap: 14px;
    padding: 14px 16px;
    box-sizing: border-box;
    border-bottom: 1px solid #e2e2dd;
    background: #ffffff;
  }

  .proj-row:last-child {
    border-bottom: none;
  }

  /* 封面 */
  .cover {
    width: 72px;
    height: 72px;
    border-radius: 4px;
    object-fit: cover;
    border: 1px solid #e2e2dd;
    background: #fbfbf9;
  }

  /* 信息区 */
  .info {
    min-width: 0;
    display: flex;
    flex-direction: column;
    gap: 6px;
  }

  .line1 {
    display: flex;
    align-items: center;
    gap: 10px;
    min-width: 0;
  }

  .tname {
    color: #2e2e2e;
    font-size: 13px;
    font-weight: 700;
    white-space: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;
  }

  .badge {
    flex: 0 0 auto;
    height: 20px;
    padding: 0 10px;
    border-radius: 100px;
    background: #ffffff;
    border: 1px solid #e2e2dd;
    color: #5f5f5f;
    font-size: 12px;
    display: inline-flex;
    align-items: center;
  }

  .oname {
    color: #2e2e2e;
    font-size: 12px;
  }

  .desc {
    color: #2e2e2e;
    font-size: 12px;
    line-height: 1.4;
    display: -webkit-box;
    -webkit-line-clamp: 3;
    -webkit-box-orient: vertical;
    overflow: hidden;
  }

  @media (max-width: 980px) {
    .page {
      grid-template-columns: 1fr;
    }
    .splitter {
      display: none;
    }
    .content {
      padding-left: 0;
    }
  }
</style>
