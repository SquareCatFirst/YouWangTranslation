<script lang="ts">
  export let value: number = 0;       // 0 ~ 100
  export let height = 6;              // 轨道高度（px）
  export let showText = true;         // 是否显示百分比文字
  export let textWidth = 36;          // 文字区域宽度（px）
  export let ariaLabel = '进度条';

  const clamp01 = (v: number) => Math.min(100, Math.max(0, Number.isFinite(v) ? v : 0));
  $: v = clamp01(value);
</script>

<div class="progress" aria-label={ariaLabel}>
  <div class="track" style={`height:${height}px`}>
    <div class="fill" style={`width:${v}%`} />
  </div>

  {#if showText}
    <div class="text" style={`width:${textWidth}px`}>{v}%</div>
  {/if}
</div>

<style>
  .progress {
    display: flex;
    align-items: center;
    gap: 10px;
    min-width: 0;
  }

  .track {
    flex: 1 1 auto;
    border-radius: 999px;
    background: #cfcfc9;
    overflow: hidden;
    position: relative;
  }

  .fill {
    height: 100%;
    background: #7f8c84; /* 你原型偏灰绿 */
    border-radius: 999px;
  }

  .text {
    flex: 0 0 auto;
    font-size: 12px;
    color: #5f5f5f;
    text-align: right;
    line-height: 1;
    white-space: nowrap;
  }
</style>
