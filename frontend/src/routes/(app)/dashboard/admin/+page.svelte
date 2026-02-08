<script lang="ts">
  import PageHeader from '$lib/components/PageHeader.svelte';
  import ProgressBar from '$lib/components/ProgressBar.svelte';

  // ===== 模板数据 =====
  type DispatchTask = {
    id: string;
    coverUrl: string;
    projectName: string;
    assignText: string;
    progress: number;
  };

  const dispatchTasks: DispatchTask[] = Array.from({ length: 10 }).map((_, i) => ({
    id: `d-${i + 1}`,
    coverUrl: `https://picsum.photos/seed/dispatch-${i + 1}/160/160`,
    projectName: i % 3 === 0 ? '间谍过家家第三季' : '桃园暗鬼',
    assignText: i % 3 === 0 ? '【千棠】校对 第六章' : '【千棠】翻译 7.1',
    progress: [15, 35, 60, 80, 25, 45, 70, 55, 90, 40][i]
  }));

  type ProjectUpdate = { id: string; text: string };
  type ProjectOverview = {
    id: string;
    coverUrl: string;
    projectName: string;
    rings: [number, number, number, number]; // 翻译/校对/嵌字/监工
    updates: ProjectUpdate[];
  };

  const overviewProjects: ProjectOverview[] = Array.from({ length: 10 }).map((_, i) => {
    const tl = [80, 95, 60, 45, 70, 85, 100, 55, 90, 40][i];
    const pr = [65, 80, 55, 30, 60, 75, 90, 40, 70, 25][i];
    const ts = [40, 70, 35, 20, 45, 55, 80, 30, 60, 15][i];
    const su = [55, 85, 45, 25, 50, 65, 88, 35, 72, 18][i];

    return {
      id: `p-${i + 1}`,
      coverUrl: `https://picsum.photos/seed/project-${i + 1}/220/220`,
      projectName:
        i % 4 === 0
          ? '桃园暗鬼'
          : i % 4 === 1
            ? '间谍过家家第三季'
            : i % 4 === 2
              ? '最新勇者与吊车尾'
              : '朋友家的妹妹真烦',
      rings: [tl, pr, ts, su],
      updates: [
        { id: `u-${i}-1`, text: '【千棠】翻译：7.1 距今 52天' },
        { id: `u-${i}-2`, text: '【工鸣蛇】校对：7.1 距今 18天' },
        { id: `u-${i}-3`, text: '【XX】嵌字：7.1 距今 6天' }
      ]
    };
  });

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

  // ===== 四环进度条（SVG 计算）=====
  const ringColors = ['#9FB8B2', '#B6B8A4', '#CDBFA6', '#CDBFA6'] as const;

  const clamp = (v: number) => Math.min(100, Math.max(0, Number.isFinite(v) ? v : 0));

  function totalProgress(rings: [number, number, number, number]) {
    const s = rings.reduce((a, b) => a + clamp(b), 0);
    return Math.round(s / 4);
  }

  function circumference(r: number) {
    return 2 * Math.PI * r;
  }

  // 从顶部开始的百分比终点坐标（用于 badge）
  function endPoint(cx: number, cy: number, r: number, percent: number) {
    const p = clamp(percent) / 100;
    const angle = -Math.PI / 2 + 2 * Math.PI * p;
    return { x: cx + r * Math.cos(angle), y: cy + r * Math.sin(angle) };
  }

  // 给模板用：返回一个“渲染需要的所有常量”
  function buildRingsModel(rings: [number, number, number, number]) {
    const size = 120;
    const cx = size / 2;
    const cy = size / 2;

    const rList = [30, 38, 46, 54];
    const strokeW = 6;

    const bg = rList.map((r) => ({ r, c: circumference(r) }));
    const fg = rList.map((r, i) => {
      const c = circumference(r);
      const v = clamp(rings[i]);
      return { r, c, v, dashA: (c * v) / 100, dashB: c };
    });

    const badges = rList.map((r, i) => {
      const pt = endPoint(cx, cy, r, rings[i]);
      return {
        x: pt.x,
        y: pt.y,
        value: clamp(rings[i]),
        color: ringColors[i]
      };
    });

    return {
      size,
      cx,
      cy,
      strokeW,
      bg,
      fg,
      total: totalProgress(rings),
      badges
    };
  }
</script>

<PageHeader title="管理员面板" />

<div class="page">
  <section class="layout">
    <!-- 左：派发任务（20%） -->
    <div class="panel">
      <div class="panel-title">派发任务</div>
      <div class="panel-divider" />

      <div class="dispatch-list">
        {#each dispatchTasks as t (t.id)}
          <div class="task-card">
            <div class="task-left">
              <img class="task-cover" src={t.coverUrl} alt={t.projectName} />
              <div class="task-name" title={t.projectName}>{t.projectName}</div>
            </div>

            <div class="task-right">
              <div class="task-assign">{t.assignText}</div>
              <ProgressBar value={t.progress} />
            </div>
          </div>
        {/each}
      </div>
    </div>

    <!-- 中：项目概览（60%） -->
    <div class="panel">
      <div class="panel-title-row">
        <div class="panel-title">项目概览</div>

        <div class="hint">
          <div class="qmark">?</div>
          <div class="tooltip">四环从里到外：翻译 / 校对 / 嵌字 / 监工。中心为总进度（均值）。</div>
        </div>
      </div>

      <div class="panel-divider" />

      <div class="overview-grid">
        {#each overviewProjects as p (p.id)}
          {@const model = buildRingsModel(p.rings)}
          <div class="overview-card">
            <!-- 左：封面+项目名 -->
            <div class="oc-left">
              <img class="oc-cover" src={p.coverUrl} alt={p.projectName} />
              <div class="oc-name" title={p.projectName}>{p.projectName}</div>
            </div>

            <!-- 中：四环进度 -->
            <div class="oc-mid">
              <div class="rings-wrap" style={`width:${model.size}px;height:${model.size}px`}>
                <svg
                  class="rings"
                  width={model.size}
                  height={model.size}
                  viewBox={`0 0 ${model.size} ${model.size}`}
                  aria-label="四环进度条"
                >
                  <g transform={`rotate(-90 ${model.cx} ${model.cy})`}>
                    {#each model.bg as b (b.r)}
                      <circle
                        class="bg"
                        cx={model.cx}
                        cy={model.cy}
                        r={b.r}
                        stroke-width={model.strokeW}
                      />
                    {/each}

                    {#each model.fg as f, idx (f.r)}
                      <circle
                        cx={model.cx}
                        cy={model.cy}
                        r={f.r}
                        stroke={ringColors[idx]}
                        stroke-width={model.strokeW}
                        stroke-linecap="round"
                        fill="none"
                        stroke-dasharray={`${f.dashA} ${f.dashB}`}
                      />
                    {/each}
                  </g>
                </svg>

                <div class="ring-center">{model.total}%</div>

                {#each model.badges as b, idx (idx)}
                  <div
                    class="ring-badge"
                    style={`left:${b.x}px; top:${b.y}px; background:${b.color}`}
                  >
                    {b.value}%
                  </div>
                {/each}
              </div>
            </div>

            <!-- 右：最新进度 -->
            <div class="oc-right">
              <div class="oc-right-title">最新进度</div>
              <div class="updates">
                {#each p.updates.slice(0, 3) as u (u.id)}
                  <div class="update-card">{u.text}</div>
                {/each}
              </div>
            </div>
          </div>
        {/each}
      </div>
    </div>

    <!-- 右：消息中心（20%） -->
    <div class="panel panel-stretch">
      <div class="panel-title">消息中心</div>
      <div class="panel-divider" />

      <div class="msg-list">
        {#each messages as m (m.id)}
          <div class="msg-card">{m.text}</div>
        {/each}
      </div>
    </div>
  </section>
</div>

<style>
  .page {
    padding: 14px 30px 40px;
  }

  /* 20 / 60 / 20 */
  .layout {
    display: grid;
    grid-template-columns: 20% 60% 20%;
    gap: 14px;
    align-items: stretch;
    min-height: 520px;
  }

  .panel {
    background: #ffffff;
    border: 1px solid #e2e2dd;
    border-radius: 6px;
    padding: 12px;
    box-sizing: border-box;
    min-width: 0;
  }

  .panel-stretch {
    display: flex;
    flex-direction: column;
    height: 100%;
  }

  .panel-title {
    font-size: 14px;
    font-weight: 600;
    color: #2e2e2e;
    padding: 2px 4px 8px;
  }

  .panel-title-row {
    display: flex;
    align-items: center;
    gap: 8px;
    padding: 2px 4px 8px;
  }

  .panel-divider {
    height: 1px;
    background: #cfcfc9;
    margin: 0 2px 12px;
  }

  /* 问号提示 */
  .hint {
    position: relative;
    display: inline-flex;
    align-items: center;
  }
  .qmark {
    width: 18px;
    height: 18px;
    border-radius: 999px;
    background: #fffcee;
    border: 1px solid #797979;
    color: #2e2e2e;
    display: grid;
    place-items: center;
    font-size: 12px;
    line-height: 1;
    cursor: default;
    user-select: none;
  }
  .tooltip {
    position: absolute;
    top: 26px;
    left: 0;
    z-index: 10;
    width: 260px;
    background: #ffffff;
    border: 1px solid #e2e2dd;
    border-radius: 6px;
    padding: 10px;
    color: #5f5f5f;
    font-size: 12px;
    line-height: 1.5;
    box-shadow: 0 10px 22px rgba(0, 0, 0, 0.06);
    display: none;
  }
  .hint:hover .tooltip {
    display: block;
  }

  /* 左：派发任务 */
  .dispatch-list {
    display: flex;
    flex-direction: column;
    gap: 10px;
    overflow: auto;
    max-height: calc(100vh - 220px);
    padding-right: 2px;
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

  /* 中：项目概览 */
  .overview-grid {
    display: grid;
    grid-template-columns: repeat(2, minmax(0, 1fr));
    gap: 12px;
    padding: 2px;
    overflow: auto;
    max-height: calc(100vh - 220px);
  }

  .overview-card {
    background: #ffffff;
    border: 1px solid #cfcfc9;
    border-radius: 6px;
    padding: 10px;
    display: grid;
    grid-template-columns: 140px 150px 1fr;
    gap: 12px;
    align-items: center;
    min-width: 0;
  }

  .oc-left {
    display: grid;
    grid-template-rows: 86px auto;
    gap: 8px;
    justify-items: start;
  }

  .oc-cover {
    width: 86px;
    height: 86px;
    border-radius: 4px;
    object-fit: cover;
    border: 1px solid #e2e2dd;
    background: #fff;
    display: block;
  }

  .oc-name {
    font-size: 12px;
    color: #2e2e2e;
    font-weight: 600;
    white-space: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;
    max-width: 140px;
  }

  .oc-mid {
    display: grid;
    place-items: center;
  }

  .rings-wrap {
    position: relative;
    display: grid;
    place-items: center;
  }

  .rings {
    display: block;
  }

  .bg {
    fill: none;
    stroke: #e2e2dd;
    opacity: 0.85;
  }

  .ring-center {
    position: absolute;
    color: #5f5f5f;
    font-size: 16px;
    font-weight: 700;
  }

  .ring-badge {
    position: absolute;
    transform: translate(-50%, -50%);
    color: #2e2e2e;
    font-size: 10px;
    font-weight: 700;
    border: 1px solid #ffffff;
    border-radius: 999px;
    padding: 2px 5px;
    line-height: 1;
    white-space: nowrap;
  }

  .oc-right {
    display: grid;
    gap: 8px;
    align-content: start;
    min-width: 0;
  }

  .oc-right-title {
    color: #2e2e2e;
    font-size: 12px;
    font-weight: 700;
  }

  .updates {
    display: grid;
    gap: 8px;
  }

  .update-card {
    background: #ffffff;
    border: 1px solid #e2e2dd;
    border-radius: 6px;
    padding: 8px 10px;
    color: #5f5f5f;
    font-size: 12px;
    line-height: 1.45;
    white-space: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;
  }

  /* 右：消息中心 */
  .msg-list {
    display: flex;
    flex-direction: column;
    gap: 10px;
    padding: 2px;
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

  @media (max-width: 1200px) {
    .layout {
      grid-template-columns: 1fr;
    }
    .overview-grid {
      grid-template-columns: 1fr;
    }
  }
</style>
