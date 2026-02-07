<script lang="ts">
  let {
    title = '',
    showSearch = false,
    searchValue = '',
    searchPlaceholder = '搜索',
    onSearch = (_q: string) => {}
  } = $props<{
    title?: string;
    showSearch?: boolean;
    searchValue?: string;
    searchPlaceholder?: string;
    onSearch?: (q: string) => void;
  }>();

  function handleInput(e: Event) {
    searchValue = (e.currentTarget as HTMLInputElement).value;
  }

  function submit() {
    onSearch(searchValue);
  }

  function handleKeydown(e: KeyboardEvent) {
    if (e.key === 'Enter') submit();
  }
</script>

<div class="page-header">
  <!-- 左 -->
  <div class="left">
    <img
      src="/material-symbols--keyboard-double-arrow-right.svg"
      alt=""
      aria-hidden="true"
      class="nav-icon"
    />
    <span class="title">{title}</span>
  </div>

  <!-- 中：搜索框 -->
  {#if showSearch}
    <div class="center">
      <div class="search-box">
        <input
          class="search-input"
          type="search"
          placeholder={searchPlaceholder}
          value={searchValue}
          oninput={handleInput}
          onkeydown={handleKeydown}
        />

        <button
          class="search-btn"
          type="button"
          aria-label="搜索"
          onclick={submit}
        >
          <img
            src="/material-symbols--search.svg"
            alt=""
            aria-hidden="true"
          />
        </button>
      </div>
    </div>
  {/if}

  <!-- 右 -->
  <div class="right">
    <slot name="actions" />
  </div>
</div>

<style>
  .page-header {
    height: 48px;
    display: grid;
    grid-template-columns: auto 1fr auto;
    align-items: center;

    background: #FBFBF9;
    border-bottom: 1px solid #E2E2DD;
    padding: 0 20px;
    box-sizing: border-box;
    gap: 16px;
  }

  /* 左 */
  .left {
    display: flex;
    align-items: center;
    gap: 8px;
  }

  .nav-icon {
    width: 16px;
    height: 16px;
    opacity: 0.7;
  }

  .title {
    font-size: 14px;
    font-weight: 600;
    color: #333;
    white-space: nowrap;
  }

  /* 中 */
  .center {
    display: flex;
    justify-content: center;
  }

  .search-box {
    display: flex;
    align-items: center;

    width: 420px;            /* ✅ 中间定宽，好看也稳定 */
    height: 35px;

    background: #FFFFFF;
    border: 1px solid #E2E2DD;
    border-radius: 10px;
    padding: 0 6px 0 10px;
    box-sizing: border-box;
  }

  .search-input {
    flex: 1;
    border: none;
    outline: none;
    background: transparent;

    font-size: 12px;
    color: #333;
  }

  .search-btn {
    width: 24px;
    height: 24px;

    border: none;
    background: transparent;
    cursor: pointer;

    display: flex;
    align-items: center;
    justify-content: center;
  }

  .search-btn img {
    width: 16px;
    height: 16px;
    opacity: 0.7;
  }

  /* 右 */
  .right {
    display: flex;
    align-items: center;
    gap: 8px;
  }
</style>
