<script lang="ts">
	import { onMount } from 'svelte';
	import Sidebar from '$lib/components/Sidebar.svelte';
	import Breadcrumb from '$lib/components/Breadcrumb.svelte';
	import favicon from '$lib/assets/favicon.svg';

	let { children } = $props();
	let role = $state(2); // 默认普通用户
	let userName = $state('123'); // 默认用户名
	let userAvatar = $state('http://ywtrans6.squarecatcloud.top:9080/i/2026/01/26/69771aba377a8.jpg'); // 默认头像
	let showSidebar = $state(false); // 控制侧边栏显示

	onMount(() => {
		const storedRole = localStorage.getItem('role');
		if (storedRole) {
			role = parseInt(storedRole);
		}
		const storedUserName = localStorage.getItem('nickname');
		if (storedUserName) {
			userName = storedUserName;
		}
		const storedAvatar = localStorage.getItem('avatarUrl');
		if (storedAvatar) {
			userAvatar = storedAvatar;
		}
	});

	type MenuItem = {
		label: string;
		href: string;
		icon?: string;
	};

	// 超级管理员菜单 (role 0)
	const superAdminMenu: MenuItem[] = [
		{ label: '首页', href: '/' },
		{ label: '用户管理', href: '/users' },
		{ label: '系统设置', href: '/settings' },
		{ label: '项目管理', href: '/projects' },
	];

	// 管理员菜单 (role 1)
	const adminMenu: MenuItem[] = [
		{ label: '首页', href: '/' },
		{ label: '用户管理', href: '/users' },
		{ label: '项目管理', href: '/projects' },
		{ label: '新建项目', href: '/new-project', icon: '/material-symbols--edit-document-outline.svg' },
	];

	// 普通用户菜单 (role 2)
	const userMenu: MenuItem[] = [
		{ label: '首页', href: '/' },
		{ label: '个人资料', href: '/profile' },
		{ label: '我的项目', href: '/my-projects' },
	];

	// 根据角色选择菜单
	let menu = $derived(() => role === 0 ? superAdminMenu : role === 1 ? adminMenu : userMenu);

	// 面包屑示例（可根据页面动态调整）
	const breadcrumb: MenuItem[] = [
		{ label: '新建项目', href: '/' , icon: '/material-symbols--edit-document-outline.svg'},
		{ label: '发布公告', href: '/users' ,icon: '/material-symbols--edit-notifications-outline.svg'},
		{ label: '消息', href: '/projects',icon: '/ic--outline-email.svg' },
	];
</script>

<svelte:head>
	<link rel="icon" href={favicon} />
</svelte:head>

<div class="top-shell">
	<div
		class="menu-trigger"
		role="button"            
		tabindex="0"            
		onmouseenter={() => (showSidebar = true)}  
		onmouseleave={() => (showSidebar = false)}
	>
		<img src="/ic--baseline-dehaze.svg" alt="Menu" class="menu-icon" />

		<Sidebar {menu} open={showSidebar} />
	</div>

	<Breadcrumb items={breadcrumb} user={{ name: userName, avatar: userAvatar }} />
</div>



<!-- 主内容区：只负责让出顶部（面包屑） -->
<div class="main-content" class:with-sidebar={showSidebar}>
	{@render children()}
</div>



<style>

	:global(:root) {
		--breadcrumb-h: 60px;
		--sidebar-w: 260px;
	}

	.top-shell {
		position: fixed;
		top: 0;
		left: 0;
		right: 0;
		height: var(--breadcrumb-h);
		z-index: 4000;
		padding: 20px;
		display: flex;              /* ✅ 修改：让 icon 和 breadcrumb 横向排列 */
		align-items: center;
		box-sizing: border-box;
		background: #FBFBF9;        /* 可选：避免下面内容透出来 */
		border-bottom: 1px solid #E2E2DD;
	}

	.menu-trigger {
		display: inline-flex;
		align-items: center;
		justify-content: center;

		width: 80px;          /* ✅ 扩大判定区域 */
		height: 80px;         /* ✅ 扩大判定区域 */

	}

	/* ✅ 修改 8：菜单 icon 样式 */
	.menu-icon {
		width: 30px;
		height: 30px;
		cursor: pointer;
	}

	.main-content {
		padding-top: var(--breadcrumb-h);
		box-sizing: border-box;
		min-height: 100vh;
		flex: 1;
		display: flex;
		flex-direction: column;
	}
</style>
