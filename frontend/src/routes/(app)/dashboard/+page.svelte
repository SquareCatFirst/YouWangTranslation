<script lang="ts">
  import PageHeader from '$lib/components/PageHeader.svelte';

  import ProgressBar from '$lib/components/ProgressBar.svelte';


  // ====== 模板数据（后续接接口直接替换）======
  type Overview = {
    todoCount: number;
    lastUpdateTitle: string;
    lastUpdateDate: string;
    mostPigeonTitle: string;
    mostPigeonInfo: string;
  };

  const overview = {
    todoCount: 14,
    lastUpdateTitle: '间谍过家家第三季 7.1',
    lastUpdateDate: '2026-1-16',
    mostPigeonTitle: '桃园暗鬼',
    mostPigeonInfo: '52天前更新'
  } satisfies Overview;

  type TaskItem = {
    id: string;
    coverUrl: string;
    projectName: string;
    assignText: string; // 例如：【千棠】翻译 7.1
    progress: number; // 0-100
  };

  const tasks: TaskItem[] = Array.from({ length: 10 }).map((_, i) => ({
    id: `t-${i + 1}`,
    coverUrl: `https://picsum.photos/seed/task-${i + 1}/160/160`,
    projectName: i % 3 === 0 ? '间谍过家家第三季' : '桃园暗鬼',
    assignText: i % 3 === 0 ? '【千棠】校对 第六章' : '【千棠】翻译 7.1',
    progress: [15, 35, 60, 80, 25, 45, 70, 55, 90, 40][i]
  }));

  type MsgItem = { id: string; text: string };

  const messages: MsgItem[] = Array.from({ length: 10 }).map((_, i) => ({
    id: `m-${i + 1}`,
    text:
      i % 3 === 0
        ? '【千棠】想要翻译《六花魔女与吊车尾勇者》'
        : i % 3 === 1
          ? '【千棠】想要翻译《间谍过家家第三季》'
          : '【千棠】翻译《六花魔女与吊车尾勇者》：已经 6 个月没有更新了'
  }));

  // ====== 进度条组件（页面内组件写法）======
  // 说明：这里用简单 div 实现，后续你也可以抽成独立组件文件
  function clamp01(v: number) {
    if (v < 0) return 0;
    if (v > 100) return 100;
    return v;
  }
</script>

<PageHeader title="个人面板" />

<div class="page">
  <section class="layout">
    <!-- 左侧大块：概览 + 我的任务 -->
    <div class="left">
      <!-- 概览 -->
      <div class="panel">
        <div class="panel-title">概览</div>
        <div class="panel-divider" />

        <div class="overview-grid">
          <div class="overview-item">
            <div class="overview-top big">{overview.todoCount}</div>
            <div class="overview-bottom">待办项目</div>
          </div>

          <div class="overview-item">
            <div class="overview-top">
              <div>{overview.lastUpdateTitle}</div>
              <div>{overview.lastUpdateDate}</div>
            </div>
            <div class="overview-bottom">上次更新</div>
          </div>

          <div class="overview-item">
            <div class="overview-top">
              <div>{overview.mostPigeonTitle}</div>
              <div>{overview.mostPigeonInfo}</div>
            </div>
            <div class="overview-bottom">最鸽项目</div>
          </div>
        </div>
      </div>

      <!-- 我的任务 -->
      <div class="panel">
        <div class="panel-title">我的任务</div>
        <div class="panel-divider" />

        <div class="task-grid">
          {#each tasks as t (t.id)}
            <div class="task-card">
              <div class="task-left">
                <img class="task-cover" src={t.coverUrl} alt={t.projectName} />
                <div class="task-name" title={t.projectName}>{t.projectName}</div>
              </div>

              <div class="task-right">
                <div class="task-assign">{t.assignText}</div>

                <!-- 进度条 -->
                <ProgressBar value={t.progress} />
              </div>
            </div>
          {/each}
        </div>
      </div>
    </div>

    <!-- 右侧大块：消息中心 -->
    <div class="right">
      <div class="panel panel-stretch">
        <div class="panel-title">消息中心</div>
        <div class="panel-divider" />

        <div class="msg-list">
          {#each messages as m (m.id)}
            <div class="msg-card">
              {m.text}
            </div>
          {/each}
        </div>
      </div>
    </div>
  </section>
</div>

<style>
  .page {
    padding: 18px 300px 300px;
  }

  /* ===== 大布局：左(概览+任务) + 右(消息中心) ===== */
  .layout {
    display: grid;
    grid-template-columns: 1fr 360px; /* 右侧像原型更窄一点 */
    gap: 14px;
    align-items: stretch;
  }

  .left {
    display: flex;
    flex-direction: column;
    gap: 14px;
    min-width: 0;
  }

  .right {
    min-width: 0;
  }

  /* ===== 大块通用（白底 + E2E2DD 边框） ===== */
  .panel {
    background: #ffffff;
    border: 1px solid #e2e2dd;
    border-radius: 6px;
    padding: 12px;
    box-sizing: border-box;
  }

  .panel-stretch {
    height: 100%;
    display: flex;
    flex-direction: column;
  }

  .panel-title {
    font-size: 14px;
    font-weight: 600;
    color: #2e2e2e;
    padding: 2px 4px 8px;
  }

  .panel-divider {
    height: 1px;
    background: #cfcfc9;
    margin: 0 2px 12px;
  }

  /* ===== 概览区域：三块横向 ===== */
  .overview-grid {
    display: grid;
    grid-template-columns: repeat(3, minmax(0, 1fr));
    gap: 10px;
    padding: 4px 2px 2px;
  }

  .overview-item {
    display: grid;
    place-items: center;
    gap: 8px;
    padding: 12px 10px;
    border-radius: 6px;
  }

  .overview-top {
    color: #2e2e2e;
    font-size: 12px;
    text-align: center;
    line-height: 1.35;
  }

  .overview-top.big {
    font-size: 28px;
    font-weight: 700;
  }

  .overview-bottom {
    color: #2e2e2e;
    font-size: 12px;
    font-weight: 600;
    text-align: center;
  }

  /* ===== 我的任务：二维平铺 ===== */
  .task-grid {
    display: grid;
    grid-template-columns: repeat(2, minmax(0, 1fr));
    gap: 12px;
    padding: 2px;
  }

  .task-card {
    background: #ffffff;
    border: 1px solid #e2e2dd;
    border-radius: 6px;
    padding: 10px;
    display: grid;
    grid-template-columns: 92px 1fr;
    gap: 12px;
    align-items: center;
    min-width: 0;
  }

  .task-left {
    display: grid;
    grid-template-rows: 64px auto;
    gap: 8px;
    justify-items: start;
  }

  .task-cover {
    width: 64px;
    height: 64px;
    border-radius: 4px;
    object-fit: cover;
    border: 1px solid #e2e2dd;
    background: #fff;
    display: block;
  }

  .task-name {
    font-size: 12px;
    color: #2e2e2e;
    white-space: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;
    max-width: 92px;
  }

  .task-right {
    display: grid;
    gap: 10px;
    min-width: 0;
  }

  .task-assign {
    color: #5f5f5f;
    font-size: 12px;
    font-weight: 600;
    white-space: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;
  }

  /* ===== 进度条组件样式 ===== */
  .progress {
    display: flex;
    align-items: center;
    gap: 10px;
    min-width: 0;
  }

  .progress-track {
    flex: 1 1 auto;
    height: 6px;
    border-radius: 999px;
    background: #cfcfc9;
    overflow: hidden;
    position: relative;
  }

  .progress-fill {
    height: 100%;
    background: #7f8c84; /* 类似原型偏灰绿 */
    border-radius: 999px;
  }

  .progress-text {
    flex: 0 0 auto;
    font-size: 12px;
    color: #5f5f5f;
    width: 36px;
    text-align: right;
  }

  /* ===== 消息中心 ===== */
  .msg-list {
    display: flex;
    flex-direction: column;
    gap: 10px;
    padding: 2px;

    /* 让消息区在右侧大块里滚动（不影响顶部标题高度） */
    flex: 1 1 auto;
    min-height: 0;
    overflow: auto;
  }

  .msg-card {
    background: #ffffff;
    border: 1px solid #e2e2dd;
    border-radius: 6px;
    padding: 10px 12px;
    color: #5f5f5f;
    font-size: 12px;
    line-height: 1.5;
  }

  /* ===== 响应式 ===== */
  @media (max-width: 1100px) {
    .layout {
      grid-template-columns: 1fr;
    }
    .task-grid {
      grid-template-columns: 1fr;
    }
  }
</style>
