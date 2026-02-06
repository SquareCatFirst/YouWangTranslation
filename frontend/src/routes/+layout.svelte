<script lang="ts">
  let { children } = $props();

  
</script>

{@render children()}

<style>
  :global(html, body) {
    margin: 0;
    padding: 0;
    background: #F6F5F1;          /* ✅ 背景色放到 body 上最稳 */
  }

  :global(:root) {
    --border-color: #E2E2DD;
  }
</style>
