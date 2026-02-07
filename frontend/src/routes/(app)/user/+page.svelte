<script lang="ts">
  import PageHeader from '$lib/components/PageHeader.svelte';
  import Pagination from '$lib/components/Pagination.svelte';

  import EnableSwitch from '$lib/components/EnableSwitch.svelte';

  let enabled = $state(true);

  // 0=超级管理员 1=管理员 2=普通用户
  type RoleLevel = 0 | 1 | 2;
  type UserStatus = '启用' | '禁用';

  type UserRow = {
    id: number;
    avatarUrl: string;
    nickname: string;
    role: RoleLevel;          // ✅ 改成数字等级
    qqOpenId: string;
    accountStatus: UserStatus;
    registerAt: string;
    lastLoginAt: string;
    enabled: boolean;
  };

  // ===== 当前登录用户角色（从 localStorage 读取）=====
  // 你登录后会做：localStorage.setItem('role', userData.role);
  let currentRole = $state<RoleLevel>(2);

  function loadCurrentRole() {
    const v = Number(localStorage.getItem('role') ?? '2');
    currentRole = (v === 0 || v === 1 || v === 2 ? v : 2) as RoleLevel;
  }

  // 首次加载读一次（runes 模式用 effect）
  $effect(() => {
    loadCurrentRole();
  });

  function roleLabel(r: RoleLevel) {
    if (r === 0) return '超级管理员';
    if (r === 1) return '管理员';
    return '普通用户';
  }

  // ✅ 权限变更规则：
  // - 只能“高等级”改“低等级”
  // - 且不能把别人改成比自己更高（比如管理员不能把别人改成超级管理员）
  // - 超级管理员(0)：可改 1/2（也可以保留 0 给自己/特殊用；这里允许改成 0 但你也可禁掉）
  // - 管理员(1)：只能改普通用户(2)（可改成 1 或 2）
  // - 普通用户(2)：不能改任何人
  function canEditRole(target: UserRow): boolean {
    if (currentRole === 2) return false;
    if (currentRole === 0) return target.role > 0;     // 超级管理员只能改低于自己（1/2）
    // currentRole === 1
    return target.role === 2;                          // 管理员只改普通用户
  }

  function roleOptionsFor(target: UserRow): RoleLevel[] {
    if (!canEditRole(target)) return [target.role];

    if (currentRole === 0) {
      // 超级管理员可以把管理员/普通用户互相切换
      return [1, 2];
    }

    // 管理员：只能把普通用户改成管理员/普通用户
    return [1, 2];
  }

  // ✅ 启用/禁用按钮规则：
  // - 超级管理员：可控制全部
  // - 管理员：只能控制普通用户
  // - 普通用户：不能控制任何人
  function canToggleEnable(target: UserRow): boolean {
    if (currentRole === 0) return true;
    if (currentRole === 1) return target.role === 2;
    return false;
  }

  // ===== 页面状态 =====
  let keyword = $state('');
  let page = $state(1);     // ✅ 1-based
  let perPage = $state(20);

  // ===== 假数据（你后续接接口时替换这里）=====
  let allUsers = $state<UserRow[]>(
    Array.from({ length: 500 }).map((_, i) => {
      const id = 1001 + i;
      const enabled = i % 4 !== 0;
      const role: RoleLevel = (i % 9 === 0 ? 0 : i % 7 === 0 ? 1 : 2) as RoleLevel;

      return {
        id,
        avatarUrl: 'https://api.dicebear.com/7.x/identicon/svg?seed=' + id,
        nickname: '工藤空',
        role,
        qqOpenId: 'aaabbbccc',
        accountStatus: enabled ? '启用' : '禁用',
        registerAt: '2026-1-13 15:01',
        lastLoginAt: '2026-1-13 15:01',
        enabled
      };
    })
  );

  // ===== 搜索（本地过滤占位）=====
  function search(v: string) {
    keyword = v;
    page = 1;
  }

  function handlePageChange(p: number) {
    page = p;
  }

  function handlePerPageChange(n: number) {
    perPage = n;
    page = 1; // ✅ 1-based
  }

  function toggleEnable(userId: number) {
    const target = allUsers.find((x) => x.id === userId);
    if (!target) return;
    if (!canToggleEnable(target)) return;

    allUsers = allUsers.map((u) => {
      if (u.id !== userId) return u;
      const nextEnabled = !u.enabled;
      return {
        ...u,
        enabled: nextEnabled,
        accountStatus: nextEnabled ? '启用' : '禁用'
      };
    });
  }

  function changeRole(userId: number, next: RoleLevel) {
    const target = allUsers.find((x) => x.id === userId);
    if (!target) return;
    if (!canEditRole(target)) return;

    // 防止赋予比自己更高的权限（数值越小权限越大）
    if (next < currentRole) return;

    allUsers = allUsers.map((u) => (u.id === userId ? { ...u, role: next } : u));
  }

  // ===== 计算：过滤后的列表 + 总数 + 当前页数据 =====
  let filtered = $state<UserRow[]>([]);
  let total = $state(0);
  let pageUsers = $state<UserRow[]>([]);

  function recompute() {
    const q = keyword.trim().toLowerCase();

    const list = allUsers.filter((u) => {
      if (!q) return true;
      return (
        String(u.id).includes(q) ||
        u.nickname.toLowerCase().includes(q) ||
        roleLabel(u.role).toLowerCase().includes(q) ||
        u.qqOpenId.toLowerCase().includes(q) ||
        u.accountStatus.toLowerCase().includes(q) ||
        u.registerAt.toLowerCase().includes(q) ||
        u.lastLoginAt.toLowerCase().includes(q)
      );
    });

    filtered = list;
    total = list.length;

    const start = (page - 1) * perPage;
    const end = start + perPage;
    pageUsers = list.slice(start, end);
  }

  $effect(() => {
    allUsers;
    keyword;
    page;
    perPage;
    currentRole;
    recompute();
  });
</script>



<PageHeader
  title="用户信息"
  showSearch={true}
  searchValue={keyword}
  searchPlaceholder="查询用户ID 昵称 权限 账号状态 注册时间 或 最近登录时间"
  onSearch={search}
/>

<div class="panel-wrap">
  <div class="panel">
    <!-- 表头 -->
    <div class="thead">
      <div class="th id">用户ID</div>
      <div class="th avatar">头像</div>
      <div class="th nick">昵称</div>
      <div class="th role">权限</div>
      <div class="th openid">QQ OpenID</div>
      <div class="th status">账号状态</div>
      <div class="th reg">注册时间</div>
      <div class="th login">最近登录时间</div>
      <div class="th enable">启用 / 禁用</div>
    </div>

    <div class="divider"></div>

    <!-- 表体 -->
    {#if pageUsers.length === 0}
      <div class="empty">暂无数据</div>
    {:else}
      <div class="tbody">
        {#each pageUsers as u (u.id)}
          <div class="tr">
            <div class="td id">{u.id}</div>

            <div class="td avatar">
              <div class="avatar-wrap">
                <img src={u.avatarUrl} alt="" />
              </div>
            </div>

            <div class="td nick">{u.nickname}</div>

            <!-- ✅ 权限：下拉框 + icon + 权限控制 -->
            <div class="td role">
              <div class="select-wrap" class:disabled={!canEditRole(u)}>
                <select
                  class="select"
                  value={u.role}
                  disabled={!canEditRole(u)}
                  onchange={(e) => changeRole(u.id, Number((e.currentTarget as HTMLSelectElement).value) as RoleLevel)}
                  aria-label="修改权限"
                >
                  {#each roleOptionsFor(u) as r (r)}
                    <option value={r}>{roleLabel(r)}</option>
                  {/each}
                </select>

                <img
                  class="dropdown-icon"
                  src="/material-symbols--arrow-drop-down-rounded.svg"
                  alt=""
                  aria-hidden="true"
                />
              </div>
            </div>

            <div class="td openid">{u.qqOpenId}</div>
            <div class="td status">{u.accountStatus}</div>
            <div class="td reg">{u.registerAt}</div>
            <div class="td login">{u.lastLoginAt}</div>

            <!-- ✅ 启用/禁用：按权限禁用 -->
            <div class="td enable">
                <EnableSwitch
                    value={u.enabled}
                    onLabel="启用"
                    offLabel="禁用"
                    disabled={!canToggleEnable(u)}
                    onChange={() => toggleEnable(u.id)}
                />
                </div>

          </div>
        {/each}
      </div>
    {/if}

    <div class="pager-wrap">
      <Pagination
        total={total}
        page={page}
        perPage={perPage}
        perPageOptions={[10, 20, 50, 100]}
        onPageChange={handlePageChange}
        onPerPageChange={handlePerPageChange}
      />
    </div>
  </div>
</div>

<style>
  .panel-wrap {
    display: flex;
    justify-content: center;
    width: 100%;
    box-sizing: border-box;
    padding: 0 24px;
  }

  .panel {
    width: 100%;
    max-width: 1200px;

    background: #ffffff;
    border: 1px solid #e2e2dd;
    border-radius: 10px;
    overflow: hidden;
    margin: 12px 0;
    box-sizing: border-box;
  }

  .divider {
    height: 1px;
    background: #cfcfc9;
  }

  .thead,
  .tr {
    display: grid;
    grid-template-columns:
      80px
      70px
      110px
      140px
      140px
      90px
      150px
      150px
      140px;
    align-items: center;
  }

  /* ✅ 表头居中 */
  .thead {
    height: 44px;
    padding: 0 16px;
    box-sizing: border-box;
    font-size: 12px;
    color: #2e2e2e;
    font-weight: 600;
  }

  .tbody .tr {
    padding: 10px 16px;
    box-sizing: border-box;
    font-size: 12px;
    color: #2e2e2e;
    border-top: 1px solid rgba(226, 226, 221, 0.55);
  }

  .tbody .tr:first-child {
    border-top: none;
  }

  /* ✅ 表头/表项全部居中 */
  .td,
  .th {
    display: flex;
    align-items: center;
    justify-content: center;  /* ✅ 水平居中 */
    text-align: center;       /* ✅ 文本居中 */
    min-width: 0;
  }

  .td.openid,
  .td.reg,
  .td.login {
    white-space: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;
  }

  .avatar-wrap {
    width: 28px;
    height: 28px;
    border-radius: 999px;
    border: 1px solid #e2e2dd;
    overflow: hidden;
    background: #f6f5f1;
    display: flex;
    align-items: center;
    justify-content: center;
  }

  .avatar-wrap img {
    width: 100%;
    height: 100%;
    object-fit: cover;
  }

  /* ✅ 下拉框（权限） */
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
    text-align: center;
  }

  .select-wrap.disabled .select,
  .select:disabled {
    cursor: default;
    opacity: 0.55;
    background: #fbfbf9;
  }

  .dropdown-icon {
    position: absolute;
    right: 6px;
    width: 16px;
    height: 16px;
    pointer-events: none;
    opacity: 0.8;
  }

  /* ✅ 启用/禁用开关 */
  .switch {
    height: 22px;
    border-radius: 999px;
    border: 1px solid #e2e2dd;
    background: #ffffff;
    display: inline-flex;
    align-items: center;
    gap: 6px;
    padding: 2px 8px 2px 2px;
    cursor: pointer;
    box-sizing: border-box;
  }

  .switch:disabled {
    cursor: default;
    opacity: 0.45;
  }

  .knob {
    width: 16px;
    height: 16px;
    border-radius: 999px;
    background: #cfcfc9;
    position: relative;
    transition: transform 0.15s ease;
  }

  .switch[aria-checked="true"] .knob {
    background: #6b6b6b;
    transform: translateX(12px);
  }

  .switch-text {
    font-size: 11px;
    color: #2e2e2e;
    user-select: none;
    white-space: nowrap;
  }

  .empty {
    padding: 18px 16px;
    font-size: 12px;
    color: #777;
    text-align: center;
  }

  .pager-wrap {
    display: flex;
    justify-content: flex-end;
    padding: 10px 10px;
    border-top: 1px solid #cfcfc9;
    background: #ffffff;
  }
</style>
