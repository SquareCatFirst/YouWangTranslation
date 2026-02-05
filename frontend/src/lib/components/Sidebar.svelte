<script lang="ts">
  import { goto } from '$app/navigation';

  let { menu, open = false } = $props();
</script>

<div class="sidebar" class:open={open}>
  <ul class="nav">
    {#each menu as item}
      <li class="nav-item">
        <a class="nav-link" href={item.href} on:click|preventDefault={() => goto(item.href)}>
          {#if item.icon}
            <img class="nav-icon" src={item.icon} alt={item.label} />
          {:else}
            <span class="nav-icon-placeholder" aria-hidden="true"></span>
          {/if}
          <span class="nav-text">{item.label}</span>
        </a>
      </li>
    {/each}
  </ul>
</div>


<style>
   
  .sidebar {
    width: var(--sidebar-w);
    background: #FBFBF9;
    padding: 20px;
    box-sizing: border-box;

    position: fixed;
    left: 0;
    top: var(--breadcrumb-h);
    height: calc(100vh - var(--breadcrumb-h));

    /* 边框 */
    border: 1px solid #CFCFC9;
    border-left: none; /* 左边贴屏幕，一般不需要边框 */

    /* 右上 & 右下圆角 */
    border-top-right-radius: 8px;
    border-bottom-right-radius: 8px;

    z-index: 2000;
    overflow: auto;

    /* 阴影轻一点，更像悬浮面板 */
    box-shadow: 2px 0 6px rgba(0, 0, 0, 0.06);

    /* ✅ 动画核心 */
    transform: translateX(-100%);
    transition: transform 0.5s ease;

    pointer-events: none;
  }

    /* ✅ 展开状态 */
  .sidebar.open {
    transform: translateX(0);
    pointer-events: auto;
  }

  ul {
    list-style: none;
    padding: 0;
  }
  li {
    margin-bottom: 10px;
  }
  a {
    text-decoration: none;
    color: #333;
    display: block;
    padding: 10px;
    border-radius: 4px;
  }
  a:hover {
    background: #ddd;
  }
  .nav {
  list-style: none;
  padding: 0;
  margin: 0;
  display: flex;
  flex-direction: column; /* ✅ 纵向排列 */
  gap: 8px;
}

  .nav-link {
    display: flex;              /* ✅ 左右排列 */
    align-items: center;
    gap: 10px;

    padding: 10px 12px;
    border-radius: 8px;
    text-decoration: none;
    color: #333;
  }

  .nav-link:hover {
    background: rgba(0, 0, 0, 0.05);
  }

  .nav-icon {
    width: 20px;
    height: 20px;
    flex: 0 0 20px;
  }

  .nav-icon-placeholder {
    width: 20px;
    height: 20px;
    flex: 0 0 20px;
  }

  .nav-text {
    font-size: 14px;
    line-height: 1;
  }

</style>