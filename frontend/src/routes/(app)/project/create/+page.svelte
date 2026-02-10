<script lang="ts">
  import PageHeader from '$lib/components/PageHeader.svelte';

  // ====== Types ======
  type ProjectCategory = '漫画' | '小说' | '视频' | '音频' | '其他';

  type ProjectInfo = {
    coverUrl: string; // 本地预览（不上传）
    name: string; // 项目名称
    alias: string; // 项目译名
    author: string; // 作者
    categories: ProjectCategory[]; // 分类（用逗号输入）
    sourceSite: string; // 连载网站
    introCn: string; // 项目简介
    introOther: string; // 翻译简介
  };

  // ====== State ======
  let project = $state<ProjectInfo>({
    coverUrl: '',
    name: '',
    alias: '',
    author: '',
    categories: [],
    sourceSite: '',
    introCn: '',
    introOther: ''
  });

  function onPickCover(e: Event) {
    const input = e.currentTarget as HTMLInputElement;
    const file = input.files?.[0];
    if (!file) return;

    // 仅本地预览（不上传）
    const url = URL.createObjectURL(file);
    project.coverUrl = url;

    // 允许同一个文件反复选择触发 change
    input.value = '';
  }

  function onClickNew() {
    console.log('新建(占位)');
  }

  function onClickExport() {
    console.log('导出(占位)', { project });
  }
</script>

<PageHeader title="新建项目">
  <div slot="actions" class="page-actions">
    <button class="primary" onclick={onClickNew}>新建</button>
    <button class="ghost" onclick={onClickExport}>导出</button>
  </div>
</PageHeader>

<div class="page">
  <div class="grid">
    <!-- ========== 项目信息 ========== -->
    <section class="card">
      <div class="card-title">项目信息</div>
      <div class="card-divider"></div>

      <!-- 上传封面 -->
      <div class="block cover-block">
        <div class="label">上传封面</div>

        <div class="cover-box">
          {#if project.coverUrl}
            <img class="cover-img" src={project.coverUrl} alt="封面预览" />
          {:else}
            <div class="cover-placeholder">请上传封面</div>
          {/if}
        </div>

        <div class="cover-actions">
          <label class="upload-btn">
            上传封面
            <input type="file" accept="image/*" onchange={onPickCover} />
          </label>

          {#if project.coverUrl}
            <button class="ghost small" type="button" onclick={() => (project.coverUrl = '')}>
              清除
            </button>
          {/if}
        </div>
      </div>

      <!-- 项目名称 + 项目译名 -->
      <div class="block">
        <div class="row">
          <div class="label">项目名称</div>
          <input class="input" bind:value={project.name} placeholder="请输入文本" />
        </div>

        <div class="row">
          <div class="label">项目译名</div>
          <input class="input" bind:value={project.alias} placeholder="请输入文本" />
        </div>
      </div>

      <!-- 作者 -->
      <div class="block">
        <div class="row">
          <div class="label">作者</div>
          <input class="input" bind:value={project.author} placeholder="请输入文本" />
        </div>
      </div>

      <!-- 分类（逗号分隔输入） -->
      <div class="block">
        <div class="row">
          <div class="label">分类</div>
          <input
            class="input"
            value={project.categories.join(',')}
            oninput={(e) => {
              const v = (e.currentTarget as HTMLInputElement).value;
              project.categories = v
                .split(',')
                .map((s) => s.trim())
                .filter(Boolean) as ProjectCategory[];
            }}
            placeholder="例如：漫画,小说"
          />
        </div>
      </div>

      <!-- 连载网站 -->
      <div class="block">
        <div class="row">
          <div class="label">连载网站</div>
          <input class="input" bind:value={project.sourceSite} placeholder="请输入文本" />
        </div>
      </div>

      <!-- 项目简介 + 翻译简介 -->
      <div class="block">
        <div class="row textarea-row">
          <div class="label">项目简介</div>
          <textarea class="textarea" bind:value={project.introCn} rows="4" placeholder="请输入文本" />
        </div>

        <div class="row textarea-row">
          <div class="label">翻译简介</div>
          <textarea
            class="textarea"
            bind:value={project.introOther}
            rows="4"
            placeholder="请输入文本"
          />
        </div>
      </div>
    </section>
  </div>
</div>

<style>
  /* 顶部右侧按钮（仅影响本页） */
  .page-actions button {
    height: 28px;
    padding: 0 12px;
    border-radius: 6px;
    font-size: 12px;
    cursor: pointer;
  }
  .primary {
    background: #333;
    color: #fff;
    border: none;
  }
  .ghost {
    background: transparent;
    color: #333;
    border: 1px solid #e2e2dd;
  }
  .small {
    height: 26px;
    padding: 0 10px;
    font-size: 12px;
  }

  .page {
    padding: 16px;
    box-sizing: border-box;
  }

  /* ✅ 只保留一列，并居中 */
  .grid {
    display: grid;
    grid-template-columns: 460px;
    justify-content: center;
    align-items: start;
    gap: 30px;
    margin: 0 auto;
    padding: 16px;
    box-sizing: border-box;
  }

  @media (max-width: 520px) {
    .grid {
      grid-template-columns: 1fr;
      padding: 0;
    }
  }

  /* 卡片 */
  .card {
    background: #fbfbf9;
    border: 1px solid #e2e2dd;
    border-radius: 10px;
    padding: 14px;
    box-sizing: border-box;
  }

  .card-title {
    font-size: 13px;
    font-weight: 700;
    color: #333;
  }

  .card-divider {
    height: 1px;
    background: #cfcfc9;
    margin: 10px 0 12px;
  }

  /* 每个板块 */
  .block {
    border: 1px solid #e2e2dd;
    background: #fbfbf9;
    border-radius: 10px;
    padding: 10px;
    box-sizing: border-box;
    margin-bottom: 12px;
  }

  /* 行：label + input 横向 */
  .row {
    display: flex;
    align-items: center;
    gap: 10px;
  }
  .row + .row {
    margin-top: 10px;
  }

  .label {
    width: 72px;
    font-size: 12px;
    color: #555;
    flex-shrink: 0;
  }

  .input,
  .textarea {
    flex: 1;
    border: 1px solid #e2e2dd;
    background: #f6f5f1;
    border-radius: 8px;
    padding: 8px 10px;
    box-sizing: border-box;
    outline: none;
    font-size: 12px;
    color: #333;
  }

  .textarea-row {
    align-items: flex-start;
  }

  .textarea {
    resize: vertical;
  }

  /* 封面上传块 */
  .cover-block {
    display: flex;
    align-items: center;
    gap: 12px;
  }

  .cover-box {
    width: 120px;
    height: 160px;
    border: 1px dashed #e2e2dd;
    border-radius: 10px;
    background: #f6f5f1;
    display: flex;
    align-items: center;
    justify-content: center;
    overflow: hidden;
    flex-shrink: 0;
  }

  .cover-img {
    width: 100%;
    height: 100%;
    object-fit: cover;
  }

  .cover-placeholder {
    font-size: 12px;
    color: #777;
  }

  .cover-actions {
    display: flex;
    flex-direction: column;
    gap: 8px;
  }

  .upload-btn {
    display: inline-flex;
    align-items: center;
    justify-content: center;
    height: 28px;
    padding: 0 12px;
    border-radius: 6px;
    border: 1px solid #e2e2dd;
    background: #fbfbf9;
    cursor: pointer;
    font-size: 12px;
    color: #333;
    user-select: none;
  }

  .upload-btn input {
    display: none;
  }
</style>
