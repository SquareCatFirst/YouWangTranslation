<script lang="ts">
  import { goto } from '$app/navigation';

  let { menu, open = false } = $props();
</script>

<div class="sidebar" class:open={open}>
  <ul>
    {#each menu as item}
      <li>
        <a href={item.href} on:click|preventDefault={() => goto(item.href)}>
          <div class="menu-item">
            {#if item.icon}
              <img src={item.icon} alt={item.label} />
            {/if}
            <span>{item.label}</span>
          </div>
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
  .menu-item {
    display: flex;
    flex-direction: column;
    align-items: center;
    text-align: center;
  }
  .menu-item img {
    width: 24px;
    height: 24px;
    margin-bottom: 4px;
  }
  .menu-item span {
    font-size: 12px;
  }
</style>