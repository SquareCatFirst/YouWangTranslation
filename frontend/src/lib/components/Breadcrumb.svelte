<script lang="ts">
  import { goto } from '$app/navigation';
  import { createEventDispatcher } from 'svelte';

  const dispatch = createEventDispatcher();
  let { items, user } = $props();
</script>

<div class="breadcrumb">
  <div class="left-icons">
    <img src="/ic--baseline-horizontal-rule.svg" alt="Separator" class="separator-icon" />
    <img src="/YwIcon.png" alt="Yw Icon" class="yw-icon" />
    <strong class="brand-name">攸望</strong>
  </div>
  <div class="breadcrumb-items">
    {#each items as item, index}
      {#if index > 0} {/if}
      <div class="breadcrumb-item">
        <a href={item.href} on:click|preventDefault={() => goto(item.href)}>
          {#if item.icon}
            <img src={item.icon} alt={item.label} />
          {/if}
          <span>{item.label}</span>
        </a>
      </div>
    {/each}
    {#if user}
      <div class="user-info">
        <span>{user.name}<br>欢迎你</span>
        {#if user.avatar}
          <img src={user.avatar} alt="用户头像" class="user-avatar" />
        {/if}
      </div>
    {/if}
  </div>
</div>

<style>
  .breadcrumb {
    width: 100%;
    flex: 1;          
    min-width: 0;    
    height: var(--breadcrumb-h);
    box-sizing: border-box;

    padding: 10px 20px;
    background: transparent; 
    border-bottom: none;     
    display: flex;
    align-items: center;

    /* ✅ 删除 z-index */
  }
  .left-icons {
    display: flex;
    align-items: center;
  }
  .left-icons svg,
  .left-icons .yw-icon,
  .left-icons .menu-icon,
  .left-icons .separator-icon,
  .left-icons .brand-name {
    margin-right: 10px;
  }
  .yw-icon,
  .menu-icon,
  .separator-icon {
    width: 30px;
    height: 30px;
  }
  .brand-name {
    font-size: 20px;
    color: #333;
  }
  .separator-icon {
    transform: rotate(90deg);
  }
  .breadcrumb-items {
    margin-left: auto;
    display: flex;
    align-items: center;
    gap: 35px;
  }
  .breadcrumb-item {
    display: flex;
    flex-direction: column;
    align-items: center;
    text-align: center;
  }
  .breadcrumb-item a {
    text-decoration: none;
    color: inherit;
    display: flex;
    flex-direction: column;
    align-items: center;
  }
  .breadcrumb-item img {
    width: 24px;
    height: 24px;
  }
  .breadcrumb-item span {
    color: #666;
  }
  .user-info {
    display: flex;
    align-items: center;
  }
  .user-info span {
    margin-right: 10px;
    color: #333;
    text-align: right;
  }
  .user-avatar {
    width: 30px;
    height: 30px;
    border-radius: 50%;
    border: 2px solid #fff;
    box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  }
</style>