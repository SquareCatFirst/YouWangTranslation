<script lang="ts">
  let { title = '' } = $props();
</script>

<div class="page-header">
  <div class="left">
    <!-- 左侧 svg -->
    <img
      src="/material-symbols--keyboard-double-arrow-right.svg"
      alt=""
      aria-hidden="true"
      class="nav-icon"
    />

    <!-- 文字直接在右边 -->
    <span class="title">{title}</span>
  </div>

  <div class="right">
    <slot name="actions" />
  </div>
</div>

<style>
  .page-header {
    height: 48px;
    display: flex;
    align-items: center;
    justify-content: space-between;

    background: #FBFBF9;
    border-bottom: 1px solid #E2E2DD;
    padding: 0 20px;
    box-sizing: border-box;
  }

  .left {
    display: flex;
    align-items: center;
    gap: 8px;                 /* svg 与文字间距 */
  }

  .nav-icon {
    width: 16px;
    height: 16px;
    opacity: 0.7;
    flex-shrink: 0;
  }

  .title {
    font-size: 14px;
    color: #333;
    font-weight: 600;
    white-space: nowrap;
  }

  .right {
    display: flex;
    align-items: center;
    gap: 8px;
  }
</style>
