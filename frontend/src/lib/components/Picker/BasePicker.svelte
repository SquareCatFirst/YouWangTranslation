<script lang="ts">
  import { tick } from 'svelte';

  export type PickerItem = { id: string };

  let {
    items = [] as PickerItem[],
    selected = [] as PickerItem[],
    onAdd = (item: PickerItem) => {},
    onRemove = (id: string) => {},

    // ✅ 渲染策略：不同数据自己定义
    getLabel = (i: PickerItem) => (i as any).name ?? (i as any).label ?? i.id,
    getSubLabel = (_i: PickerItem) => '',
    getAvatar = (_i: PickerItem) => '',

    // ✅ 搜索策略（两选一/可同时提供）
    filterFn = (i: PickerItem, q: string) => getLabel(i).toLowerCase().includes(q),
    onSearch, // (q:string)=>Promise<void>|void

    // ✅ 行为
    single = false,
    placeholder = '搜索...'
  } = $props<{
    items: PickerItem[];
    selected: PickerItem[];
    onAdd?: (item: PickerItem) => void;
    onRemove?: (id: string) => void;

    getLabel?: (item: PickerItem) => string;
    getSubLabel?: (item: PickerItem) => string;
    getAvatar?: (item: PickerItem) => string;

    filterFn?: (item: PickerItem, query: string) => boolean;
    onSearch?: (query: string) => void | Promise<void>;

    single?: boolean;
    placeholder?: string;
  }>();

  let open = $state(false);
  let query = $state('');
  let searchEl = $state<HTMLInputElement | null>(null);

  // ✅ 模板安全数组
  let selectedSafe = $state<PickerItem[]>([]);
  let results = $state<PickerItem[]>([]);

  function recompute() {
    const itemsArr = Array.isArray(items) ? items : [];
    const selectedArr = Array.isArray(selected) ? selected : [];
    selectedSafe = selectedArr;

    const q = query.trim().toLowerCase();
    const selectedIds = new Set(selectedArr.map((s) => s.id));

    // 如果提供了 onSearch（远程），BasePicker 不做过滤，直接展示 items（但仍排除已选）
    // 如果没提供 onSearch（本地），BasePicker 用 filterFn 过滤
    const base = itemsArr.filter((i) => !selectedIds.has(i.id));

    results = onSearch
      ? base.slice(0, 50)
      : base
          .filter((i) => (!q ? true : filterFn(i, q)))
          .slice(0, 50);
  }

  $effect(() => {
    items;
    selected;
    query;
    recompute();
  });

  async function openDialog() {
    open = true;
    query = '';
    await tick();
    searchEl?.focus();
    // 打开时也触发一次远程搜索（可选）
    if (onSearch) await onSearch('');
  }

  function close() {
    open = false;
    query = '';
  }

  function onKeydown(e: KeyboardEvent) {
    if (e.key === 'Escape') close();
  }

  async function handleQueryInput(v: string) {
    query = v;
    if (onSearch) {
      await onSearch(v.trim());
    }
  }

  function addItem(i: PickerItem) {
    onAdd(i);
    if (single) close();
    query = '';
    searchEl?.focus();
  }
</script>

<div class="picker">
  <div class="chips">
    {#each selectedSafe as s (s.id)}
      <div class="chip">
        <span class="chip-name">{getLabel(s)}</span>
        <button class="chip-x" type="button" onclick={() => onRemove(s.id)} aria-label="移除">✕</button>
      </div>
    {/each}
  </div>

  <button class="add-btn" type="button" onclick={openDialog} aria-label="添加">
    <img class="add-icon" src="/material-symbols--add-circle-outline.svg" alt="" aria-hidden="true" />
  </button>

  {#if open}
    <div class="overlay" role="presentation" onkeydown={onKeydown}>
      <button class="backdrop" type="button" aria-label="关闭弹窗" onclick={close}></button>

      <div class="dialog" role="dialog" aria-modal="true" aria-label="搜索">
        <div class="dialog-head">
          <div class="dialog-title">搜索</div>
          <button class="close-btn" type="button" onclick={close} aria-label="关闭">✕</button>
        </div>

        <div class="search">
          <input
            class="search-input"
            placeholder={placeholder}
            bind:this={searchEl}
            value={query}
            oninput={(e) => handleQueryInput((e.currentTarget as HTMLInputElement).value)}
          />
        </div>

        <div class="list">
          {#if results.length === 0}
            <div class="empty">没有匹配结果</div>
          {:else}
            {#each results as r (r.id)}
              <button class="row" type="button" onclick={() => addItem(r)}>
                <div class="row-left">
                  {#if getAvatar(r)}
                    <img class="avatar" src={getAvatar(r)} alt="" />
                  {:else}
                    <div class="avatar placeholder" aria-hidden="true">{getLabel(r).slice(0, 1)}</div>
                  {/if}

                  <div class="meta">
                    <div class="name">{getLabel(r)}</div>
                    {#if getSubLabel(r)}
                      <div class="email">{getSubLabel(r)}</div>
                    {/if}
                  </div>
                </div>

                <div class="row-right">添加</div>
              </button>
            {/each}
          {/if}
        </div>
      </div>
    </div>
  {/if}
</div>

<style>
  .picker {
    display: flex;
    align-items: center;
    gap: 8px;
    width: 100%;

    background: #f5f5f5;
    border: 1px solid #e2e2dd;
    border-radius: 10px;
    padding: 6px 8px;
    box-sizing: border-box;
  }

  .chips {
    display: flex;
    align-items: center;
    gap: 8px;
    flex-wrap: wrap;
    flex: 1;
    min-width: 0;
  }

  .chip {
    display: inline-flex;
    align-items: center;
    gap: 6px;
    height: 26px;
    padding: 0 10px;
    border-radius: 10px;
    border: 1px solid #e2e2dd;
    background: #fbfbf9;
    font-size: 12px;
    color: #333;
  }

  .chip-x {
    width: 18px;
    height: 18px;
    border-radius: 6px;
    border: 1px solid #e2e2dd;
    background: transparent;
    cursor: pointer;
    font-size: 12px;
    line-height: 16px;
    padding: 0;
    color: #666;
  }

  .add-btn {
    flex-shrink: 0;
    width: 28px;
    height: 28px;
    padding: 0;
    border: none;
    background: transparent;
    cursor: pointer;
    display: inline-flex;
    align-items: center;
    justify-content: center;
  }

  .add-icon { width: 20px; height: 20px; opacity: 0.9; }

  .overlay {
    position: fixed;
    inset: 0;
    z-index: 9999;
    display: flex;
    align-items: center;
    justify-content: center;
    padding: 16px;
    box-sizing: border-box;
  }

  .backdrop {
    position: fixed;
    inset: 0;
    border: none;
    padding: 0;
    margin: 0;
    background: rgba(0, 0, 0, 0.18);
    cursor: pointer;
  }

  .dialog {
    position: relative;
    width: min(520px, 100%);
    background: #fbfbf9;
    border: 1px solid #e2e2dd;
    border-radius: 12px;
    box-shadow: 0 12px 40px rgba(0, 0, 0, 0.18);
    overflow: hidden;
    z-index: 1;
  }

  .dialog-head {
    display: flex;
    align-items: center;
    justify-content: space-between;
    padding: 12px 14px;
    border-bottom: 1px solid #e2e2dd;
  }

  .dialog-title { font-size: 13px; font-weight: 700; color: #333; }

  .close-btn {
    width: 28px;
    height: 28px;
    border-radius: 8px;
    border: 1px solid #e2e2dd;
    background: transparent;
    cursor: pointer;
    color: #666;
  }

  .search { padding: 12px 14px; border-bottom: 1px solid #e2e2dd; }

  .search-input {
    width: 100%;
    border: 1px solid #e2e2dd;
    background: #f6f5f1;
    border-radius: 10px;
    padding: 10px 12px;
    box-sizing: border-box;
    font-size: 12px;
    outline: none;
  }

  .list { max-height: 320px; overflow: auto; }

  .row {
    width: 100%;
    border: none;
    background: transparent;
    cursor: pointer;
    padding: 10px 14px;
    display: flex;
    align-items: center;
    justify-content: space-between;
    text-align: left;
  }

  .row:hover { background: rgba(0, 0, 0, 0.04); }

  .row-left { display: flex; align-items: center; gap: 10px; min-width: 0; }

  .avatar {
    width: 32px;
    height: 32px;
    border-radius: 10px;
    object-fit: cover;
    border: 1px solid #e2e2dd;
    background: #f6f5f1;
    flex-shrink: 0;
    display: flex;
    align-items: center;
    justify-content: center;
  }

  .avatar.placeholder { font-size: 12px; color: #555; }

  .meta { min-width: 0; }

  .name {
    font-size: 12px;
    font-weight: 600;
    color: #333;
    white-space: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;
  }

  .email {
    font-size: 12px;
    color: #777;
    white-space: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;
  }

  .row-right {
    font-size: 12px;
    color: #333;
    border: 1px solid #e2e2dd;
    background: #fbfbf9;
    border-radius: 10px;
    padding: 6px 10px;
    flex-shrink: 0;
  }

  .empty { padding: 18px 14px; font-size: 12px; color: #777; }
</style>
