<script lang="ts">
  let {
    total = 0,
    page = 1, // ✅ 1-based
    perPage = 20,
    perPageOptions = [10, 20, 50, 100],
    onPageChange = (_p: number) => {},
    onPerPageChange = (_n: number) => {}
  } = $props<{
    total: number;
    page: number; // ✅ 1-based
    perPage: number;
    perPageOptions?: number[];
    onPageChange?: (p: number) => void;      // ✅ 回传 1-based
    onPerPageChange?: (n: number) => void;
  }>();

  // 内部统一用 0-based 计算
  let totalPages = $state(1);
  let last0 = $state(0);
  let cur0 = $state(0);
  let prev0 = $state(0);
  let next0 = $state(0);

  function clamp(n: number, min: number, max: number) {
    return Math.max(min, Math.min(max, n));
  }

  function recompute() {
    const pp = Math.max(1, Number(perPage) || 1);
    const t = Math.max(0, Number(total) || 0);

    totalPages = Math.max(1, Math.ceil(t / pp));
    last0 = totalPages - 1;

    // ✅ 外部 page 是 1-based，这里转成 0-based
    const p0 = clamp((Number(page) || 1) - 1, 0, last0);

    cur0 = p0;
    prev0 = clamp(p0 - 1, 0, last0);
    next0 = clamp(p0 + 1, 0, last0);
  }

  $effect(() => {
    total;
    perPage;
    page;
    recompute();
  });

  // ✅ 传入显示页码（1-based），回调也用 1-based
  function go(displayP: number) {
    const target0 = clamp(displayP - 1, 0, last0);
    const target1 = target0 + 1;
    if (target1 === page) return;
    onPageChange(target1);
  }

  function goPrev() {
    if (cur0 <= 0) return;
    onPageChange(cur0); // (cur0-1)+1 = cur0
  }

  function goNext() {
    if (cur0 >= last0) return;
    onPageChange(cur0 + 2); // (cur0+1)+1
  }

  function changePerPage(e: Event) {
    const v = Number((e.currentTarget as HTMLSelectElement).value);
    onPerPageChange(v);
    onPageChange(1); // ✅ 切条数后回到第 1 页（1-based）
  }

    // ✅ 显示用 1-based（runes：用 $state + $effect）
  let displayCur = $state(1);
  let displayPrev = $state(1);
  let displayNext = $state(1);
  let displayLast = $state(1);

  $effect(() => {
    cur0;
    prev0;
    next0;
    last0;

    displayCur = cur0 + 1;
    displayPrev = prev0 + 1;
    displayNext = next0 + 1;
    displayLast = last0 + 1;
  });

</script>

<div class="pager">
  <div class="total">共 {total} 条</div>
  <div class="vline" aria-hidden="true"></div>

  <div class="perpage">
    <span>每页显示</span>

    <div class="select-wrap">
      <select class="select" value={perPage} onchange={changePerPage} aria-label="每页显示条数">
        {#each perPageOptions as n (n)}
          <option value={n}>{n}</option>
        {/each}
      </select>

      <img class="dropdown-icon" src="/material-symbols--arrow-drop-down-rounded.svg" alt="" aria-hidden="true" />
    </div>

    <span>条</span>
  </div>

  <div class="vline" aria-hidden="true"></div>

  <button class="icon-btn" type="button" onclick={goPrev} disabled={cur0 <= 0} aria-label="上一页">
    <img class="nav-icon" src="/material-symbols--arrow-back-ios-new-rounded.svg" alt="" aria-hidden="true" />
  </button>

  <div class="pages" aria-label="页码">
    <!-- ✅ prev：当前页 >= 3 才显示（避免出现 0 或 1 的前一页） -->
    {#if displayCur >= 3}
      <button class="page" type="button" onclick={() => go(displayPrev)}>
        {displayPrev}
      </button>
    {/if}

    <!-- 当前页 -->
    <button class="page current" type="button" aria-current="page" disabled>
      {displayCur}
    </button>

    <!-- ✅ next：没到最后页才显示 -->
    {#if cur0 < last0}
      <button class="page" type="button" onclick={() => go(displayNext)}>
        {displayNext}
      </button>
    {/if}

    <!-- ✅ 省略号 + 最后一页：距离最后页至少 2 页才显示 -->
    {#if cur0 < last0 - 1}
      <span class="dots">…</span>
      <button class="page" type="button" onclick={() => go(displayLast)}>
        {displayLast}
      </button>
    {/if}
  </div>

  <button class="icon-btn" type="button" onclick={goNext} disabled={cur0 >= last0} aria-label="下一页">
    <img class="nav-icon right" src="/material-symbols--arrow-back-ios-new-rounded.svg" alt="" aria-hidden="true" />
  </button>
</div>

<style>
  .pager {
    display: flex;
    align-items: center;
    gap: 10px;
    height: 40px;
    padding: 0 10px;
    box-sizing: border-box;
    background: #fff;
  }

  .total {
    font-size: 12px;
    color: #2e2e2e;
    white-space: nowrap;
  }

  .vline {
    width: 1px;
    height: 16px;
    background: #cfcfc9;
    flex-shrink: 0;
  }

  .perpage {
    display: flex;
    align-items: center;
    gap: 6px;
    font-size: 12px;
    color: #2e2e2e;
    white-space: nowrap;
  }

  .select-wrap {
    position: relative;
    display: inline-flex;
    align-items: center;
    height: 24px;
  }

  .select {
    height: 24px;
    font-size: 12px;
    color: #2e2e2e;
    padding: 0 22px 0 10px;
    border: 1px solid #e2e2dd;
    border-radius: 10px;
    background: #fff;
    outline: none;
    appearance: none;
    cursor: pointer;
  }

  .dropdown-icon {
    position: absolute;
    right: 6px;
    width: 16px;
    height: 16px;
    pointer-events: none;
    opacity: 0.8;
  }

  .icon-btn {
    width: 28px;
    height: 28px;
    border: none;
    background: transparent;
    padding: 0;
    display: inline-flex;
    align-items: center;
    justify-content: center;
    cursor: pointer;
  }

  .icon-btn:disabled {
    cursor: default;
    opacity: 0.35;
  }

  .nav-icon {
    width: 18px;
    height: 18px;
    opacity: 0.85;
  }

  .nav-icon.right {
    transform: rotate(180deg);
  }

  .pages {
    display: flex;
    align-items: center;
    gap: 8px;
  }

  .page {
    min-width: 22px;
    height: 24px;
    padding: 0 6px;
    border: none;
    background: transparent;
    font-size: 12px;
    color: #2e2e2e;
    cursor: pointer;
    border-radius: 6px;
  }

  .page:disabled {
    color: #bdbdbd;
    cursor: default;
  }

  .page.current {
    background: #f6f5f1;
    border: 1px solid #797979;
    color: #2e2e2e;
    cursor: default;
  }

  .dots {
    font-size: 12px;
    color: #2e2e2e;
    opacity: 0.7;
    user-select: none;
  }
</style>
