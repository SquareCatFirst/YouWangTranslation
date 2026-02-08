<script lang="ts">
  import EnableSwitch from '$lib/components/EnableSwitch.svelte';

  // ===== 静态资源（放在 frontend/static 下时，页面用 /xxx.svg 访问）======
  const ICON_BACK = '/material-symbols--arrow-left-alt.svg';
  const ICON_ARROW = '/material-symbols--arrow-back-ios-new-rounded.svg';
  const ICON_DROP = '/material-symbols--arrow-drop-down-rounded.svg';

  const ICON_COPY = '/material-symbols--file-copy-outline.svg';
  const ICON_LP = '/material-symbols--lab-profile.svg';
  const ICON_ADD_PHOTO = '/material-symbols--add-photo-alternate-rounded.svg';
  const ICON_CHECK = '/material-symbols--check-circle-rounded.svg';

  // ===== 顶部栏（页码）======
  let chapterName = '1.1';
  let currentPage = 1;
  const pageOptions = Array.from({ length: 20 }).map((_, i) => i + 1);

  function prevPage() {
    currentPage = Math.max(1, currentPage - 1);
  }
  function nextPage() {
    currentPage = Math.min(pageOptions.length, currentPage + 1);
  }

  // ===== 图片/舞台 =====
  let enabled = true; // 生肉/熟肉（示例）
  const imageUrl = 'https://picsum.photos/seed/workpage/1200/1600';

  let viewportEl: HTMLDivElement | null = null;
  let imgEl: HTMLImageElement | null = null;

  // 舞台尺寸（= 图片自然尺寸）
  let imgW = 1;
  let imgH = 1;

  // 变换：以舞台左上为原点：transform = translate(tx, ty) scale(scale)
  let scale = 1;
  let tx = 0;
  let ty = 0;

  // 拖拽
  let dragging = false;
  let dragStartX = 0;
  let dragStartY = 0;
  let dragStartTx = 0;
  let dragStartTy = 0;

  // 用于判断“单击 vs 拖拽”
  let moved = false;

  // 添加标号模式
  let addMode = false;

  type MarkerType = 'inside' | 'outside';
  type Marker = {
    id: number;
    x: number; // 舞台坐标（像素）
    y: number; // 舞台坐标（像素）
    type: MarkerType;
    jp: string;
    cn: string;
    history: { at: string; text: string }[];
  };

  let markers: Marker[] = [];
  let selectedId: number | null = null;

  function onImageLoad() {
    if (!imgEl) return;
    imgW = imgEl.naturalWidth || 1;
    imgH = imgEl.naturalHeight || 1;

    // 初始：按左上角顶点展示
    scale = 1;
    tx = 0;
    ty = 0;
    clampTransform();
  }

  function clampTransform() {
    if (!viewportEl) return;

    const vw = viewportEl.clientWidth;
    const vh = viewportEl.clientHeight;

    const scaledW = imgW * scale;
    const scaledH = imgH * scale;

    // 不允许露出空白
    const minTx = Math.min(0, vw - scaledW);
    const maxTx = 0;
    const minTy = Math.min(0, vh - scaledH);
    const maxTy = 0;

    tx = Math.max(minTx, Math.min(maxTx, tx));
    ty = Math.max(minTy, Math.min(maxTy, ty));
  }

  // ===== 缩放（滚轮）=====
  function onWheel(e: WheelEvent) {
    if (!viewportEl) return;

    const rect = viewportEl.getBoundingClientRect();
    const mx = e.clientX - rect.left;
    const my = e.clientY - rect.top;

    const stageX = (mx - tx) / scale;
    const stageY = (my - ty) / scale;

    const delta = e.deltaY;
    const factor = delta > 0 ? 0.92 : 1.08;

    const nextScale = Math.max(0.2, Math.min(4, scale * factor));
    if (nextScale === scale) return;

    scale = nextScale;
    tx = mx - stageX * scale;
    ty = my - stageY * scale;

    clampTransform();
  }

  // ===== 拖拽平移 =====
  function onPointerDown(e: PointerEvent) {
    moved = false;

    // addMode 下不允许拖拽（只允许点图添加标号 / 点空白退出）
    if (addMode) return;

    dragging = true;
    dragStartX = e.clientX;
    dragStartY = e.clientY;
    dragStartTx = tx;
    dragStartTy = ty;
    (e.currentTarget as HTMLElement).setPointerCapture(e.pointerId);
  }

  function onPointerMove(e: PointerEvent) {
    if (!dragging) return;

    const dx = e.clientX - dragStartX;
    const dy = e.clientY - dragStartY;

    // 超过阈值认为是拖拽，不是单击
    if (Math.abs(dx) > 3 || Math.abs(dy) > 3) moved = true;

    tx = dragStartTx + dx;
    ty = dragStartTy + dy;
    clampTransform();
  }

    function onPointerUp(e: PointerEvent) {
        // ✅ addMode 下：用这个 pointerup 去添加标号 or 点空白退出
        if (addMode) {
            tryAddMarkerByPointerUp(e);
            return;
        }

        dragging = false;
        try {
            (e.currentTarget as HTMLElement).releasePointerCapture(e.pointerId);
        } catch {}

        // 单击空白 => 取消选中
        if (!moved) {
            selectedId = null;
        }
        }



  // ===== 添加标号 =====
  // 你要的行为：点击 + 进入 addMode；再点 + 退出；点空白退出；点图片添加标号但不退出
  function enterAddModeToggle() {
    addMode = !addMode;
    if (addMode) selectedId = null; // 进入添加模式时，先取消选中
  }

  function nowStr() {
    const d = new Date();
    const pad = (n: number) => String(n).padStart(2, '0');
    return `${d.getFullYear()}/${d.getMonth() + 1}/${d.getDate()} ${pad(d.getHours())}:${pad(
      d.getMinutes()
    )}:${pad(d.getSeconds())}`;
  }

  // addMode 下：pointerup 用来判断点在图上还是空白
  function tryAddMarkerByPointerUp(e: PointerEvent) {
    if (!addMode) return;
    if (!viewportEl) return;

    const rect = viewportEl.getBoundingClientRect();
    const mx = e.clientX - rect.left;
    const my = e.clientY - rect.top;

    const stageX = (mx - tx) / scale;
    const stageY = (my - ty) / scale;

    // 点空白（不在图片范围）=> 退出 addMode
    if (stageX < 0 || stageY < 0 || stageX > imgW || stageY > imgH) {
      addMode = false;
      return;
    }

    // 点在图片上 => 添加标号（但不退出 addMode）
    const id = (markers.at(-1)?.id ?? 0) + 1;
    const now = nowStr();

    const m: Marker = {
      id,
      x: stageX,
      y: stageY,
      type: 'inside',
      jp: '',
      cn: '',
      history: [
        { at: now, text: `#${id} 千棠 创建了该标签` },
        { at: now, text: `#${id} 千棠 创建了该标签（中文/日文为空）` }
      ]
    };

    markers = [...markers, m];
    selectedId = id; // 添加后默认选中新标号，右侧立即可编辑
    // addMode 保持 true
  }

  function selectMarker(id: number) {
    selectedId = id;
    addMode = false; // 选中标号进入编辑，不再是添加模式（更符合你右侧编辑的逻辑）
  }

  function clearSelected() {
    selectedId = null;
    addMode = false;
  }

  // ===== 右侧编辑绑定 =====
  function getSelected(): Marker | null {
    if (selectedId == null) return null;
    return markers.find((m) => m.id === selectedId) ?? null;
  }

  function updateSelected(patch: Partial<Marker>) {
    if (selectedId == null) return;
    markers = markers.map((m) => (m.id === selectedId ? { ...m, ...patch } : m));
  }

  function saveSelected() {
    const m = getSelected();
    if (!m) return;
    const now = nowStr();
    const text = `#${m.id} 千棠 保存了该标签（中文/日文已更新）`;
    updateSelected({ history: [{ at: now, text }, ...m.history] });
  }

  // 操作栏显示
  $: zoomText = `${(scale * 100).toFixed(2)}%`;
  $: offsetXPct = imgW > 0 ? Math.max(0, (-tx / (imgW * scale)) * 100) : 0;
  $: offsetYPct = imgH > 0 ? Math.max(0, (-ty / (imgH * scale)) * 100) : 0;

  // 选中标号的相对坐标（0-1）
  $: selected = getSelected();
  $: selectedX = selected ? (selected.x / imgW).toFixed(3) : '0.000';
  $: selectedY = selected ? (selected.y / imgH).toFixed(3) : '0.000';

  function setType(t: MarkerType) {
    if (!selected) return;
    updateSelected({ type: t });
  }

  // viewer 光标
  $: viewerCursor = addMode ? 'crosshair' : dragging ? 'grabbing' : 'grab';
</script>

<div class="page">
  <!-- 左：工作区（70%） -->
  <section class="work">
    <!-- 顶部栏 -->
    <div class="topbar">
      <div class="tb-left">
        <button class="icon-btn" type="button" aria-label="返回章节列表">
          <img class="icon" src={ICON_BACK} alt="返回" />
        </button>
        <span class="tb-text">返回章节列表</span>
      </div>

      <div class="tb-mid">{chapterName}</div>

      <div class="tb-right">
        <button class="page-btn" type="button" on:click={prevPage}>
          <img class="icon small" src={ICON_ARROW} alt="上一页" />
          <span>上一页</span>
        </button>

        <span class="vline" />

        <div class="page-select">
          <select class="page-native" bind:value={currentPage} aria-label="页码选择">
            {#each pageOptions as p}
              <option value={p}>{p}</option>
            {/each}
          </select>
          <div class="page-box">
            <span class="page-num">{currentPage}</span>
            <img class="icon drop" src={ICON_DROP} alt="下拉" />
          </div>
        </div>

        <span class="vline" />

        <button class="page-btn" type="button" on:click={nextPage}>
          <span>下一页</span>
          <img class="icon small rotate" src={ICON_ARROW} alt="下一页" />
        </button>
      </div>
    </div>

    <!-- 图片展示区 -->
         <div
        class="viewer"
        bind:this={viewportEl}
        on:wheel|preventDefault={onWheel}
        on:pointerdown={onPointerDown}
        on:pointermove={onPointerMove}
        on:pointerup={onPointerUp}
        on:pointercancel={onPointerUp}
        style={`cursor:${viewerCursor};`}
        >

      <!-- 舞台：img + markers 同一个 transform，保证标号跟随缩放/移动 -->
      <div
        class="stage"
        style={`width:${imgW}px;height:${imgH}px;transform:translate(${tx}px,${ty}px) scale(${scale});`}
      >
        <img class="page-img" bind:this={imgEl} src={imageUrl} alt="工作图片" on:load={onImageLoad} />

        {#each markers as m (m.id)}
        <button
            type="button"
            class="marker"
            class:active={selectedId === m.id}
            style={`left:${m.x}px; top:${m.y}px;`}
            on:pointerdown|stopPropagation
            on:pointerup|stopPropagation
            on:click|stopPropagation={() => selectMarker(m.id)}
            aria-label={`标号 ${m.id}`}
        >
            {m.id}
        </button>
        {/each}

      </div>

      <!-- 右下角 + 按钮：toggle addMode -->
      <button
        class="add-btn"
        class:mode-on={addMode}
        type="button"
        on:pointerdown|stopPropagation
        on:click|stopPropagation={enterAddModeToggle}
        aria-label="添加标号"
        title={addMode ? '再次点击退出标号模式' : '点击进入标号模式'}
      >
        +
      </button>
    </div>

    <!-- 操作栏 -->
    <div class="bottombar">
      <!-- 左：图片部分 -->
      <div class="bb-left">
        <div class="zoom-offset">
          <div class="zo-row">
            <span class="zo-label">缩放</span>
            <span class="zo-split" />
            <span class="zo-value">{zoomText}</span>
          </div>
          <div class="zo-row">
            <span class="zo-label">偏移</span>
            <span class="zo-split" />
            <span class="zo-value">{offsetXPct.toFixed(0)}%, {offsetYPct.toFixed(0)}%</span>
          </div>
        </div>

        <div class="switch-wrap">
          <EnableSwitch value={enabled} onLabel="生肉" offLabel="熟肉" disabled={false} onChange={() => (enabled = !enabled)} />
        </div>

        <button class="pill-btn" type="button">
          <img class="icon" src={ICON_COPY} alt="下载" />
          <span>下载图片</span>
        </button>
      </div>

      <!-- 右：上传提交 -->
      <div class="bb-right">
        <button class="pill-btn" type="button">
          <img class="icon" src={ICON_LP} alt="上传LP" />
          <span>上传LP文档</span>
        </button>

        <button class="pill-btn" type="button">
          <img class="icon" src={ICON_ADD_PHOTO} alt="上传熟肉" />
          <span>上传熟肉</span>
        </button>

        <button class="pill-btn" type="button">
          <img class="icon" src={ICON_CHECK} alt="提交" />
          <span>提交</span>
        </button>
      </div>
    </div>
  </section>

  <!-- 右：标号编辑（30%） -->
  <aside class="editor" class:disabled={selectedId == null}>
    <!-- 顶部类型选择 -->
    <div class="type-tabs">
      <button
        class="tab"
        class:active={selected?.type === 'inside'}
        type="button"
        on:click={() => setType('inside')}
        disabled={selectedId == null}
      >
        框内
      </button>
      <div class="tab-split" />
      <button
        class="tab"
        class:active={selected?.type === 'outside'}
        type="button"
        on:click={() => setType('outside')}
        disabled={!selected}
      >
        框外
      </button>
    </div>

    <!-- 文本编辑区 -->
    <div class="edit-area">
      <div class="field">
        <div class="field-title">日文文本</div>
        <textarea
          class="ta"
          rows="4"
          placeholder="请输入日文"
          disabled={!selected}
          value={selected?.jp ?? ''}
          on:input={(e) => updateSelected({ jp: (e.target as HTMLTextAreaElement).value })}
        />
      </div>

      <div class="field">
        <div class="field-title">中文文本</div>
        <textarea
          class="ta"
          rows="4"
          placeholder="请输入中文"
          disabled={!selected}
          value={selected?.cn ?? ''}
          on:input={(e) => updateSelected({ cn: (e.target as HTMLTextAreaElement).value })}
        />
      </div>
    </div>

    <div class="hline" />

    <!-- 修改历史 -->
    <div class="history">
      <div class="history-title">修改历史</div>

      <div class="history-list">
        {#if selected}
          {#each (selected.history.length
            ? selected.history
            : Array.from({ length: 10 }).map((_, i) => ({
                at: `2026/1/11 16:32:${String(25 - i).padStart(2, '0')} #${selected.id}`,
                text: '千棠 创建了该标签'
              }))) as h, idx (idx)}
            <div class="hist-item">
              <div class="hist-top">{h.at}</div>
              <div class="hist-text">{h.text}</div>
              <div class="hist-split" />
            </div>
          {/each}
        {:else}
          {#each Array.from({ length: 10 }) as _, i (i)}
            <div class="hist-item">
              <div class="hist-top">2026/1/11 16:32:25 #1</div>
              <div class="hist-text">千棠 创建了该标签</div>
              <div class="hist-split" />
            </div>
          {/each}
        {/if}
      </div>
    </div>

    <!-- 底部操作区 -->
    <div class="editor-bottom">
      <div class="hline" />

      <div class="bottom-row">
        <div class="coords">
          <div>X: {selectedX}</div>
          <div>Y: {selectedY}</div>
        </div>

        <div class="bottom-actions">
          <button class="btn-save" type="button" on:click={saveSelected} disabled={!selected}>保存</button>
          <button class="btn-cancel" type="button" on:click={clearSelected} disabled={!selected}>取消</button>
        </div>
      </div>
    </div>
  </aside>
</div>

<style>
  .page {
    display: grid;
    grid-template-columns: 70% 30%;
    gap: 14px;
    padding: 6px 20px 0px;
    height: 92vh;
    box-sizing: border-box;
    background: #f6f5f1;
  }

  /* ================= 左侧工作区 ================= */
  .work {
    display: grid;
    grid-template-rows: 70px 1fr 70px;
    gap: 0;
    min-width: 0;
    height: 100%;
  }

  .topbar {
    height: 70px;
    background: #f6f5f1;
    border: 1px solid #e2e2dd;
    border-radius: 8px 8px 0 0;
    display: grid;
    grid-template-columns: 1fr auto 1fr;
    align-items: center;
    padding: 0 12px;
    box-sizing: border-box;
  }

  .tb-left {
    display: flex;
    align-items: center;
    gap: 8px;
    color: #2e2e2e;
  }

  .icon-btn {
    width: 32px;
    height: 32px;
    border: none;
    background: transparent;
    padding: 0;
    display: grid;
    place-items: center;
    cursor: pointer;
  }

  .tb-text {
    font-size: 14px;
  }

  .tb-mid {
    color: #2e2e2e;
    font-size: 14px;
    font-weight: 600;
    justify-self: center;
  }

  .tb-right {
    justify-self: end;
    display: flex;
    align-items: center;
    gap: 10px;
    color: #2e2e2e;
  }

  .page-btn {
    display: inline-flex;
    align-items: center;
    gap: 6px;
    border: none;
    background: transparent;
    color: #2e2e2e;
    cursor: pointer;
    font-size: 14px;
  }

  .vline {
    width: 1px;
    height: 16px;
    background: #cfcfc9;
  }

  .page-select {
    position: relative;
    display: inline-flex;
    align-items: center;
  }

  .page-native {
    position: absolute;
    inset: 0;
    opacity: 0;
    cursor: pointer;
  }

  .page-box {
    height: 28px;
    background: #ffffff;
    border: 1px solid #e2e2dd;
    border-radius: 6px;
    display: inline-flex;
    align-items: center;
    gap: 6px;
    padding: 0 8px;
    box-sizing: border-box;
    user-select: none;
  }

  .page-num {
    min-width: 10px;
    text-align: center;
    font-size: 14px;
    color: #2e2e2e;
  }

  .icon {
    width: 18px;
    height: 18px;
    display: block;
  }
  .icon.small {
    width: 16px;
    height: 16px;
  }
  .icon.rotate {
    transform: rotate(180deg);
  }
  .icon.drop {
    width: 18px;
    height: 18px;
  }

  /* 图片展示区 */
  .viewer {
    background: #ffffff;
    border-left: 1px solid #e2e2dd;
    border-right: 1px solid #e2e2dd;
    border-bottom: 1px solid #e2e2dd;
    position: relative;
    overflow: hidden;
    min-width: 0;
    touch-action: none;
  }

  .stage {
    position: absolute;
    left: 0;
    top: 0;
    transform-origin: 0 0;
  }

  .page-img {
    width: 100%;
    height: 100%;
    display: block;
    user-select: none;
    -webkit-user-drag: none;
    pointer-events: none;
  }

  .marker {
    position: absolute;
    transform: translate(-50%, -50%);
    width: 26px;
    height: 26px;
    border-radius: 999px;
    background: #ffffff00;
    border: 2px solid #7f8f86;
    color: #f90e0e;
    font-weight: 700;
    font-size: 13px;
    display: grid;
    place-items: center;
    cursor: pointer;
    user-select: none;
  }
  .marker.active {
    box-shadow: 0 0 0 3px rgba(127, 143, 134, 0.25);
  }

  .add-btn {
    position: absolute;
    right: 18px;
    bottom: 18px;
    width: 46px;
    height: 46px;
    border-radius: 999px;
    background: #ffffff;
    border: 3px solid #7f8f86;
    color: #5f5f5f;
    font-size: 26px;
    line-height: 1;
    display: grid;
    place-items: center;
    cursor: pointer;
    transition: transform 0.08s ease;
  }
  /* 添加模式下给一点视觉提示 */
  .add-btn.mode-on {
    transform: scale(1.05);
  }

  /* 底部操作栏 */
  .bottombar {
    height: 70px;
    border: 1px solid #e2e2dd;
    border-top: none;
    background: #f6f5f1;
    border-radius: 0 0 8px 8px;
    display: grid;
    grid-template-columns: 1fr 1fr;
    align-items: center;
    padding: 0 12px;
    box-sizing: border-box;
    gap: 10px;
  }

  .bb-left,
  .bb-right {
    display: flex;
    align-items: center;
    gap: 10px;
    min-width: 0;
  }
  .bb-right {
    justify-content: flex-end;
  }

  .zoom-offset {
    display: grid;
    gap: 6px;
    padding-right: 10px;
  }

  .zo-row {
    display: inline-flex;
    align-items: center;
    gap: 10px;
    font-size: 12px;
    color: #2e2e2e;
    white-space: nowrap;
  }

  .zo-label {
    min-width: 32px;
  }

  .zo-split {
    width: 1px;
    height: 12px;
    background: #cfcfc9;
  }

  .switch-wrap {
    display: flex;
    align-items: center;
  }

  .pill-btn {
    height: 34px;
    background: #ffffff;
    border: 1px solid #e2e2dd;
    border-radius: 500px;
    padding: 0 12px;
    display: inline-flex;
    align-items: center;
    gap: 8px;
    cursor: pointer;
    color: #2e2e2e;
    font-size: 12px;
    white-space: nowrap;
  }

  /* ================= 右侧编辑区 ================= */
  .editor {
    background: #f6f5f1;
    border: 1px solid #e2e2dd;
    border-radius: 8px;
    display: flex;
    flex-direction: column;
    min-width: 0;
    height: 100%;
    overflow: hidden;
  }

  .editor.disabled {
    opacity: 0.55;
    pointer-events: none; /* 禁用时不响应输入 */
  }

  .type-tabs {
    height: 60px;
    background: #ffffff;
    border-bottom: 1px solid #cfcfc9;
    display: grid;
    grid-template-columns: 1fr 1px 1fr;
    align-items: stretch;
    pointer-events: auto;
  }

  .tab {
    border: none;
    background: #ffffff;
    color: #2e2e2e;
    font-size: 14px;
    cursor: pointer;
  }
  .tab:disabled {
    cursor: not-allowed;
  }
  .tab.active {
    font-weight: 700;
  }

  .tab-split {
    background: #cfcfc9;
    width: 1px;
    height: 100%;
  }

  .edit-area {
    padding: 12px;
    display: grid;
    gap: 12px;
  }

  .field-title {
    color: #2e2e2e;
    font-size: 12px;
    margin-bottom: 6px;
  }

  .ta {
    width: 100%;
    resize: none;
    padding: 10px;
    border: 1px solid #cfcfc9;
    border-radius: 8px;
    background: #ffffff;
    font-size: 12px;
    color: #333333;
    box-sizing: border-box;
    outline: none;
  }

  .hline {
    height: 1px;
    background: #cfcfc9;
    margin: 0 12px;
  }

  .history {
    padding: 12px;
    display: flex;
    flex-direction: column;
    gap: 10px;
    min-height: 0;
    flex: 1 1 auto;
  }

  .history-title {
    color: #2e2e2e;
    font-size: 14px;
    font-weight: 700;
  }

  .history-list {
    overflow: auto;
    min-height: 0;
    padding-right: 4px;
  }

  .hist-item {
    padding: 8px 0;
  }

  .hist-top {
    color: #333333;
    font-size: 11px;
    margin-bottom: 4px;
  }

  .hist-text {
    color: #333333;
    font-size: 12px;
    margin-bottom: 8px;
  }

  .hist-split {
    height: 1px;
    background: #cfcfc9;
  }

  .editor-bottom {
    flex: 0 0 auto;
    padding: 10px 12px 12px;
    box-sizing: border-box;
    background: #f6f5f1;
  }

  .bottom-row {
    display: flex;
    align-items: flex-end;
    justify-content: space-between;
    gap: 10px;
    padding-top: 10px;
  }

  .coords {
    color: #333333;
    font-size: 12px;
    display: grid;
    gap: 6px;
  }

  .bottom-actions {
    display: flex;
    gap: 10px;
    align-items: center;
  }

  .btn-save,
  .btn-cancel {
    height: 30px;
    border-radius: 8px;
    border: 1px solid #e2e2dd;
    padding: 0 14px;
    color: #ffffff;
    cursor: pointer;
    font-size: 12px;
    pointer-events: auto;
  }

  .btn-save {
    background: #7f8f86;
  }

  .btn-cancel {
    background: #c27a7a;
  }

  @media (max-width: 1100px) {
    .page {
      grid-template-columns: 1fr;
      height: auto;
    }
    .work {
      height: 70vh;
    }
    .editor {
      height: 60vh;
    }
  }
</style>
