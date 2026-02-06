<script lang="ts">
  import BasePicker, { type PickerItem } from '$lib/components/Picker/BasePicker.svelte';

  export type User = {
    id: string;
    name: string;
    email?: string;
    avatar?: string;
  };

  let {
    users = [] as User[],
    selected = [] as User[],
    onAdd = (u: User) => {},
    onRemove = (id: string) => {},
    placeholder = '搜索用户'
  } = $props<{
    users: User[];
    selected: User[];
    onAdd?: (u: User) => void;
    onRemove?: (id: string) => void;
    placeholder?: string;
  }>();

  // ✅ 让 BasePicker 用 label 做 chips 显示
  function getLabel(item: PickerItem): string {
    return (item as User).name;
  }

  // ✅ 本地过滤（后续如果要走接口，在 BasePicker 的 onSearch 做就行）
  function filterFn(item: PickerItem, q: string): boolean {
    const u = item as User;
    const qq = q.trim().toLowerCase();
    if (!qq) return true;
    return (
      u.name.toLowerCase().includes(qq) ||
      (u.email?.toLowerCase().includes(qq) ?? false)
    );
  }

  // ✅ 处理添加：回传 User
  function handleAdd(item: PickerItem) {
    onAdd(item as User);
  }

  // ✅ 处理移除：BasePicker 传 id（string）
  function handleRemove(id: string) {
    onRemove(id);
  }
</script>

<!-- ✅ 外层用 div 包住，做你要的灰底 + 右侧 icon 固定效果（具体布局由 BasePicker 控制） -->
<div class="user-picker">
  <BasePicker
    items={users}
    selected={selected}
    single={false}
    onAdd={handleAdd}
    onRemove={handleRemove}
    getLabel={getLabel}
    filterFn={filterFn}
    placeholder={placeholder}

  />
</div>

<style>
  /* ✅ 灰底横向撑满：让父容器决定宽度（通常 row 里右侧容器需要 flex:1） */
  .user-picker :global(.picker) {
    width: 100%;
  }

  /* ✅ 如果你的 BasePicker 最外层 class 就是 .picker（和你之前一样），这里直接覆写它 */
  .user-picker :global(.picker) {
    display: flex;
    align-items: center;
    gap: 8px;

    background: #f5f5f5;
    border: 1px solid #e2e2dd;
    border-radius: 10px;

    padding: 6px 8px;
    box-sizing: border-box;
  }

  /* ✅ chips 占满，让 + 按钮永远在最右 */
  .user-picker :global(.chips) {
    flex: 1;
    min-width: 0;
  }

  .user-picker :global(.add-btn) {
    flex-shrink: 0;
  }
</style>
