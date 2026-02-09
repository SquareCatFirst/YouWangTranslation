<script lang="ts">
  import PageHeader from '$lib/components/PageHeader.svelte';

  import UserPicker, { type User } from '$lib/components/Picker/UserPicker.svelte';

  import ChapterPicker, { type Chapter } from '$lib/components/Picker/ChapterPicker.svelte';


  type ChapterImage = {
    id: string;
    fileName: string;
  };

  type ChapterInfo = {
    id: string;
    name: string;
    visible: boolean;     // ✅ bool：供展示（✓/✕）
    images: ChapterImage[];
  };

  // ✅ 章节列表（页面状态）
  let chaptersInfo = $state<ChapterInfo[]>([]);

  // ✅ 为每个章节保存一个 file input 引用（用对象 map 存）
  let fileInputs = $state<Record<string, HTMLInputElement | null>>({});

  function uid() {
    return crypto.randomUUID ? crypto.randomUUID() : String(Date.now() + Math.random());
  }

  function addChapter() {
  chaptersInfo = [
    ...chaptersInfo,
    {
      id: uid(),
      name: '',
      visible: false,
      images: [],
    }
  ];
}


  function removeChapter(id: string) {
    chaptersInfo = chaptersInfo.filter((c) => c.id !== id);
    // 清掉 input 引用
    const copy = { ...fileInputs };
    delete copy[id];
    fileInputs = copy;
  }

  function triggerPickFiles(chapterId: string) {
    // ✅ 只是触发选择文件；上传逻辑你后面写
    fileInputs[chapterId]?.click();
  }

  function onPickedFiles(chapterId: string) {
    // ✅ 你说上传逻辑后面写：这里先只演示“把文件名加入列表”
    const el = fileInputs[chapterId];
    if (!el?.files) return;

    const names = Array.from(el.files).map((f) => f.name);
    const target = chaptersInfo.find((c) => c.id === chapterId);
    if (!target) return;

    const newImgs: ChapterImage[] = names.map((n) => ({ id: uid(), fileName: n }));

    chaptersInfo = chaptersInfo.map((c) =>
      c.id === chapterId
        ? { ...c, images: [...c.images, ...newImgs] }
        : c
    );

    // 清空 input，方便同名文件也能再次选择触发 change
    el.value = '';
  }

  function removeImage(chapterId: string, imageId: string) {
    chaptersInfo = chaptersInfo.map((c) =>
      c.id === chapterId
        ? { ...c, images: c.images.filter((i) => i.id !== imageId) }
        : c
    );
  }


  // ✅ 章节结构（占位）
  type Chapter = { id: string; name: string };

  // ✅ 每条章节指派
  type ChapterAssignment = {
  id: string;
  chapterIds: string[]; // ✅ 多选章节
  users: User[];
};


  // ✅ 章节数据（先假数据，后面你接接口替换）
  let chapters = $state<Chapter[]>([
    { id: 'c1', name: '第 1 话' },
    { id: 'c2', name: '第 2 话' },
    { id: 'c3', name: '第 3 话' }
  ]);

  // ✅ 章节指派列表
  let chapterAssignments = $state<ChapterAssignment[]>([]);

  function addChapterAssignment() {
  chapterAssignments = [
    ...chapterAssignments,
    { id: crypto.randomUUID(), chapterIds: [], users: [] }
  ];
}


  function removeChapterAssignment(id: string) {
    chapterAssignments = chapterAssignments.filter((a) => a.id !== id);
  }


  function updateAssignmentChapters(assignId: string, ids: string[]) {
  chapterAssignments = chapterAssignments.map((a) =>
    a.id === assignId ? { ...a, chapterIds: ids } : a
  );
}


  function addAssignmentUser(assignmentId: string, u: User) {
    chapterAssignments = chapterAssignments.map((a) => {
      if (a.id !== assignmentId) return a;
      if (a.users.some((x) => x.id === u.id)) return a;
      return { ...a, users: [...a.users, u] };
    });
  }

  function removeAssignmentUser(assignmentId: string, userId: string) {
    chapterAssignments = chapterAssignments.map((a) => {
      if (a.id !== assignmentId) return a;
      return { ...a, users: a.users.filter((x) => x.id !== userId) };
    });
  }

  // ✅ 测试用：全量用户列表（后续你接接口就把这里换成请求结果）
  let allUsers = $state<User[]>([
    { id: 'u1', name: '千棠', email: 'qian@example.com' },
    { id: 'u2', name: '小白', email: 'bai@example.com' },
    { id: 'u3', name: '阿柒', email: 'qi@example.com' },
    { id: 'u4', name: '柚子', email: 'youzi@example.com' },
    { id: 'u5', name: '一一', email: 'yi@example.com' },
    { id: 'u6', name: '十三', email: 'shi@example.com' }
  ]);

  // ✅ 每个角色：人数限制 + 已选人员
  let adminLimit = $state<number>(2);
  let adminSelected = $state<User[]>([]);

  let sourceLimit = $state<number>(2);
  let sourceSelected = $state<User[]>([]);

  let supervisorLimit = $state<number>(2);
  let supervisorSelected = $state<User[]>([]);

  let translatorLimit = $state<number>(5);
  let translatorSelected = $state<User[]>([]);

  let proofreaderLimit = $state<number>(2);
  let proofreaderSelected = $state<User[]>([]);

  let typesetterLimit = $state<number>(2);
  let typesetterSelected = $state<User[]>([]);

  // ✅ 通用：添加/移除（带人数限制）
  function addWithLimit(selectedArr: User[], limit: number, u: User): User[] {
    if (selectedArr.some((x) => x.id === u.id)) return selectedArr;
    if (selectedArr.length >= limit) return selectedArr; // 超过限制不加（你也可以改成提示 toast）
    return [...selectedArr, u];
  }

  function removeById(selectedArr: User[], id: string): User[] {
    return selectedArr.filter((u) => u.id !== id);
  }

  // ====== Types ======
  type ProjectCategory = '漫画' | '小说' | '视频' | '音频' | '其他';

  type ProjectInfo = {
    coverUrl: string; // 预览用（本地/远端都行）
    name: string;
    alias: string;
    author: string;
    categories: ProjectCategory[];
    sourceSite: string;
    introCn: string;
    introOther: string;
  };

  type RoleKey = 'superAdmin' | 'admin' | 'editor' | 'reviewer' | 'translator' | 'viewer';

  type PermissionGroup = {
    id: string;
    label: string;
    // 简化：成员用字符串，后续你接入用户表后再改成 userId
    members: string[];
    // 每组的权限标签（你后续可以改成更细粒度）
    perms: {
      read: boolean;
      write: boolean;
      publish: boolean;
      manage: boolean;
    };
  };

  type ChapterLang = 'zh' | 'en' | 'ja' | 'ko' | 'fr' | 'de' | 'es';

  type ChapterPriceItem = {
    tier: number;   // 例如 1~11
    price: number;  // 价格（先 number 占位）
    unit: 'jp' | 'coin';
  };


  // ====== Helpers ======

  const defaultPrices = (unit: 'jp' | 'coin' = 'jp'): ChapterPriceItem[] =>
    Array.from({ length: 11 }, (_, i) => ({
      tier: i + 1,
      price: 0,
      unit
    }));

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

  // 权限：按“角色组”先占位，你后续接接口把 members / perms 替换即可
  let permissionGroups = $state<PermissionGroup[]>([
    {
      id: uid(),
      label: '项目负责人',
      members: ['默认用户A'],
      perms: { read: true, write: true, publish: true, manage: true }
    },
    {
      id: uid(),
      label: '翻译',
      members: ['默认译者1'],
      perms: { read: true, write: true, publish: false, manage: false }
    },
    {
      id: uid(),
      label: '审核',
      members: ['默认审核1'],
      perms: { read: true, write: false, publish: true, manage: false }
    }
  ]);


  // ====== UI Actions (no HTTP) ======
  function toggleCategory(cat: ProjectCategory) {
    if (project.categories.includes(cat)) {
      project.categories = project.categories.filter((c) => c !== cat);
    } else {
      project.categories = [...project.categories, cat];
    }
  }

  function addPermissionGroup() {
    permissionGroups = [
      ...permissionGroups,
      {
        id: uid(),
        label: '新权限组',
        members: [],
        perms: { read: true, write: false, publish: false, manage: false }
      }
    ];
  }

  function removePermissionGroup(id: string) {
    permissionGroups = permissionGroups.filter((g) => g.id !== id);
  }




  function onPickCover(e: Event) {
    const input = e.currentTarget as HTMLInputElement;
    const file = input.files?.[0];
    if (!file) return;

    // 仅本地预览（不上传）
    const url = URL.createObjectURL(file);
    project.coverUrl = url;
  }

  function onClickNew() {
    // 仅占位：你后续做新建逻辑
    console.log('新建(占位)');
  }

  function onClickExport() {
    // 仅占位：你后续做导出逻辑
    console.log('导出(占位)', { project, permissionGroups, chapters });
  }
</script>

<PageHeader title="修改项目">
  <div slot="actions" class="page-actions">
    <button class="primary" onclick={onClickNew}>新建</button>
    <button class="ghost" onclick={onClickExport}>导出</button>
  </div>
</PageHeader>

<div class="page">
  <div class="grid">
    <!-- ========== 项目信息 ========== -->
    <section class="card">
        <!-- 标题 -->
        <div class="card-title">项目信息</div>
        <div class="card-divider"></div>

        <!-- 上传封面（div 封装，边框 E2E2DD） -->
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
                <button class="ghost small" onclick={() => (project.coverUrl = '')}>清除</button>
            {/if}
            </div>
        </div>

        <!-- 项目名称 + 项目译名（一个 div 封装，纵向两行；每行横向 label + input） -->
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

        <!-- 作者板块（div 封装，横向 label + input） -->
        <div class="block">
            <div class="row">
            <div class="label">作者</div>
            <input class="input" bind:value={project.author} placeholder="请输入文本" />
            </div>
        </div>

        <!-- 分类板块（div 封装，右边输入框） -->
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

        <!-- 连载网站板块（div 封装，右边输入框） -->
        <div class="block">
            <div class="row">
            <div class="label">连载网站</div>
            <input class="input" bind:value={project.sourceSite} placeholder="请输入文本" />
            </div>
        </div>

        <!-- 项目简介 + 翻译简介（一个 div 封装；两行纵向；每行横向 label + textarea） -->
        <div class="block">
            <div class="row textarea-row">
            <div class="label">项目简介</div>
            <textarea class="textarea" bind:value={project.introCn} rows="4" placeholder="请输入文本"></textarea>
            </div>

            <div class="row textarea-row">
            <div class="label">翻译简介</div>
            <textarea class="textarea" bind:value={project.introOther} rows="4" placeholder="请输入文本"></textarea>
            </div>
        </div>
    </section>

    <!-- ========== 项目权限 ========== -->
   <section class="card">
  <div class="card-title">项目权限</div>
  <div class="card-divider"></div>

  <div class="fixed-title">固定指派</div>

  <!-- 管理员指派 -->
  <div class="block">
    <div class="row">
      <div class="label">人数限制</div>
      <input class="input" type="number" min="0" bind:value={adminLimit} placeholder="请输入" />
    </div>

    <div class="row">
      <div class="label">管理员</div>
      <div class="picker-wrap">
        <UserPicker
          users={allUsers}
          selected={adminSelected}
          onAdd={(u) => (adminSelected = addWithLimit(adminSelected, adminLimit, u))}
          onRemove={(id) => (adminSelected = removeById(adminSelected, id))}
        />
      </div>
    </div>
  </div>

  <!-- 图源指派 -->
  <div class="block">
    <div class="row">
      <div class="label">人数限制</div>
      <input class="input" type="number" min="0" bind:value={sourceLimit} placeholder="请输入" />
    </div>

    <div class="row">
      <div class="label">图源</div>
      <div class="picker-wrap">
        <UserPicker
          users={allUsers}
          selected={sourceSelected}
          onAdd={(u) => (sourceSelected = addWithLimit(sourceSelected, sourceLimit, u))}
          onRemove={(id) => (sourceSelected = removeById(sourceSelected, id))}
        />
      </div>
    </div>
  </div>

  <!-- 监督指派 -->
  <div class="block">
    <div class="row">
      <div class="label">人数限制</div>
      <input class="input" type="number" min="0" bind:value={supervisorLimit} placeholder="请输入" />
    </div>

    <div class="row">
      <div class="label">监督</div>
      <div class="picker-wrap">
        <UserPicker
          users={allUsers}
          selected={supervisorSelected}
          onAdd={(u) => (supervisorSelected = addWithLimit(supervisorSelected, supervisorLimit, u))}
          onRemove={(id) => (supervisorSelected = removeById(supervisorSelected, id))}
        />
      </div>
    </div>
  </div>

  <!-- 翻译指派 -->
  <div class="block">
    <div class="row">
      <div class="label">人数限制</div>
      <input class="input" type="number" min="0" bind:value={translatorLimit} placeholder="请输入" />
    </div>

    <div class="row">
      <div class="label">翻译</div>
      <div class="picker-wrap">
        <UserPicker
          users={allUsers}
          selected={translatorSelected}
          onAdd={(u) => (translatorSelected = addWithLimit(translatorSelected, translatorLimit, u))}
          onRemove={(id) => (translatorSelected = removeById(translatorSelected, id))}
        />
      </div>
    </div>
  </div>

  <!-- 校对指派 -->
  <div class="block">
    <div class="row">
      <div class="label">人数限制</div>
      <input class="input" type="number" min="0" bind:value={proofreaderLimit} placeholder="请输入" />
    </div>

    <div class="row">
      <div class="label">校对</div>
      <div class="picker-wrap">
        <UserPicker
          users={allUsers}
          selected={proofreaderSelected}
          onAdd={(u) => (proofreaderSelected = addWithLimit(proofreaderSelected, proofreaderLimit, u))}
          onRemove={(id) => (proofreaderSelected = removeById(proofreaderSelected, id))}
        />
      </div>
    </div>
  </div>

  <!-- 嵌字指派 -->
  <div class="block">
    <div class="row">
      <div class="label">人数限制</div>
      <input class="input" type="number" min="0" bind:value={typesetterLimit} placeholder="请输入" />
    </div>

    <div class="row">
      <div class="label">嵌字</div>
      <div class="picker-wrap">
        <UserPicker
          users={allUsers}
          selected={typesetterSelected}
          onAdd={(u) => (typesetterSelected = addWithLimit(typesetterSelected, typesetterLimit, u))}
          onRemove={(id) => (typesetterSelected = removeById(typesetterSelected, id))}
        />
      </div>
    </div>
  </div>

  <div class="card-divider"></div>

<!-- ✅ 章节指派：左标题 + 右按钮 -->
<div class="sub-head">
  <div class="sub-title">章节指派</div>
  <button class="ghost small" type="button" onclick={addChapterAssignment}>添加指派</button>
</div>

<!-- ✅ 动态生成每条章节指派 -->
{#if chapterAssignments.length === 0}
  <div class="placeholder">暂无章节指派，点击右上角“添加指派”创建</div>
{:else}
  {#each chapterAssignments as a (a.id)}
    <div class="block">
      <div class="row">
        <div class="label">章节</div>

        <!-- ✅ 用 ChapterPicker 替换 select（多选） -->
        <div class="picker-wrap">
          <ChapterPicker
            chapters={chapters}
            values={a.chapterIds}
            onChange={(ids) => updateAssignmentChapters(a.id, ids)}
          />
        </div>

        <button
          class="ghost small"
          type="button"
          onclick={() => removeChapterAssignment(a.id)}
          aria-label="删除该指派"
        >
          删除
        </button>
      </div>

      <div class="row">
        <div class="label">人员</div>
        <div class="picker-wrap">
          <UserPicker
            users={allUsers}
            selected={a.users}
            onAdd={(u) => addAssignmentUser(a.id, u)}
            onRemove={(id) => removeAssignmentUser(a.id, id)}
          />
        </div>
      </div>
    </div>
  {/each}
{/if}
</section>

    <!-- ========== 章节信息 ========== -->
    <!-- ========== 章节信息 ========== -->
<section class="card">
  <div class="card-title">章节信息</div>
  <div class="card-divider"></div>

  <!-- 右侧：添加章节按钮 -->
  <div class="sub-head">
    <div class="sub-title"></div>
    <button class="ghost small" type="button" onclick={addChapter}>添加章节</button>
  </div>

  {#if chaptersInfo.length === 0}
    <div class="placeholder">暂无章节，点击右上角“添加章节”创建</div>
  {:else}
    {#each chaptersInfo as ch (ch.id)}
      <div class="block">
        <!-- 第一行：章节名 + 输入框 + bool + 删除 -->
        <div class="row">
          <div class="label">章节名</div>

          <input
            class="input"
            placeholder="例如：1.1"
            bind:value={ch.name}
          />

          <!-- bool：只能 勾 / 叉 -->
          <button
            class="bool-btn"
            type="button"
            aria-label="切换供展示"
            onclick={() => (ch.visible = !ch.visible)}
            title="供展示"
          >
            {#if ch.visible}
              ✓
            {:else}
              ✕
            {/if}
          </button>

          <!-- 删除 icon -->
          <button
            class="icon-btn"
            type="button"
            aria-label="删除章节"
            onclick={() => removeChapter(ch.id)}
          >
            <img
              class="icon"
              src="/material-symbols--delete-outline.svg"
              alt=""
              aria-hidden="true"
            />
          </button>
        </div>

        <!-- 第二行：章节图 + 添加图片 icon -->
        <div class="row">
        <div class="label">章节图</div>

        <!-- ✅ 新增一层：背景容器 -->
        <div class="images-wrap">
            <div class="images">
            <!-- 添加图片 -->
            <button
                class="icon-btn add-img"
                type="button"
                aria-label="添加图片"
                onclick={() => triggerPickFiles(ch.id)}
                title="添加图片"
            >
                <img
                class="icon"
                src="/material-symbols--add-circle-outline.svg"
                alt=""
                aria-hidden="true"
                />
            </button>

            <!-- 隐藏 file input -->
            <input
                class="file-input"
                type="file"
                multiple
                accept="image/*"
                bind:this={fileInputs[ch.id]}
                onchange={() => onPickedFiles(ch.id)}
            />

            <!-- 文件名列表 -->
            {#if ch.images.length === 0}
                <div class="empty-images">暂无图片</div>
            {:else}
                <div class="file-list">
                {#each ch.images as img (img.id)}
                    <div class="file-pill">
                    <span class="file-name">{img.fileName}</span>
                    <button
                        class="pill-x"
                        type="button"
                        aria-label="删除图片"
                        onclick={() => removeImage(ch.id, img.id)}
                    >
                        ✕
                    </button>
                    </div>
                {/each}
                </div>
      {/if}
    </div>
  </div>
</div>
      </div>
    {/each}
  {/if}
</section>


  </div>
</div>

<style>

        /* 章节图整体背景 */
    .images-wrap {
    flex: 1;
    background: #F5F5F5;
    border-radius: 8px;
    padding: 10px;
    box-sizing: border-box;
    }

    /* 原 images 不需要背景了，只负责布局 */
    .images {
    display: flex;
    align-items: flex-start;
    gap: 8px;
    flex-wrap: wrap;
    }

    /* 文件列表 */
    .file-list {
    display: flex;
    flex-wrap: wrap;
    gap: 8px;
    }

    /* 空状态 */
    .empty-images {
    font-size: 12px;
    color: #888;
    }

    /* 单个文件 pill */
    .file-pill {
    display: inline-flex;
    align-items: center;
    gap: 6px;
    height: 26px;
    padding: 0 10px;
    border-radius: 10px;
    background: #FBFBF9;
    border: 1px solid #E2E2DD;
    font-size: 12px;
    }


    /* sub-head：右侧按钮那一行 */
  .sub-head {
    display: flex;
    align-items: center;
    justify-content: space-between;
    padding: 10px 0 6px;
  }

  .sub-title {
    font-size: 12px;
    color: #666;
  }

  /* 章节图区域 */
  .images {
    flex: 1;
    min-width: 0;
    display: flex;
    align-items: flex-start;
    gap: 10px;
  }

  .add-img {
    flex-shrink: 0;
  }

  .file-input {
    display: none;
  }

  .empty-images {
    font-size: 12px;
    color: #888;
    padding: 6px 0;
  }

  .file-list {
    display: flex;
    flex-wrap: wrap;
    gap: 8px;
    min-width: 0;
  }

  .file-pill {
    display: inline-flex;
    align-items: center;
    gap: 8px;

    height: 28px;
    padding: 0 10px;
    border-radius: 10px;
    border: 1px solid #e2e2dd;
    background: #fbfbf9;

    font-size: 12px;
    color: #333;
    max-width: 180px;
  }

  .file-name {
    white-space: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;
    max-width: 130px;
  }

  .pill-x {
    width: 18px;
    height: 18px;
    border-radius: 6px;
    border: 1px solid #e2e2dd;
    background: transparent;
    cursor: pointer;
    font-size: 12px;
    line-height: 16px;
    padding: 0;
    color: #666;
    flex-shrink: 0;
  }

  /* bool按钮（✓/✕） */
  .bool-btn {
    width: 32px;
    height: 28px;
    border-radius: 8px;
    border: 1px solid #e2e2dd;
    background: #fbfbf9;
    cursor: pointer;
    font-size: 14px;
    color: #333;
    flex-shrink: 0;
  }

  /* icon 按钮（删除章 / 添加图） */
  .icon-btn {
    width: 32px;
    height: 28px;
    padding: 0;
    border: 1px solid #e2e2dd;
    border-radius: 8px;
    background: #fbfbf9;
    cursor: pointer;

    display: inline-flex;
    align-items: center;
    justify-content: center;
    flex-shrink: 0;
  }

  .icon {
    width: 18px;
    height: 18px;
    opacity: 0.9;
  }

    /* ✅ 章节指派头部：左标题右按钮 */
  .sub-head {
    display: flex;
    align-items: center;
    justify-content: space-between;
    margin: 10px 0 8px;
  }

  .sub-title {
    font-size: 12px;
    font-weight: 700;
    color: #333;
  }

  /* 你已有的 placeholder 若没有，就补一个 */
  .placeholder {
    padding: 10px 0;
    font-size: 12px;
    color: #777;
  }

    /* 固定指派标题 */
  .fixed-title {
    margin: 10px 0 8px;
    font-size: 12px;
    font-weight: 700;
    color: #333;
  }

  /* 复用你已有的 block 风格：这里补齐一个常见版本（如果你已有可删） */
  .block {
    border: 1px solid #e2e2dd;
    background: #fbfbf9;
    border-radius: 10px;
    padding: 12px;
    box-sizing: border-box;
    margin-bottom: 12px;
  }

  .row {
    display: flex;
    align-items: center;
    gap: 12px;
  }

  .row + .row {
    margin-top: 10px;
  }

  .label {
    width: 72px;
    flex: 0 0 72px;
    font-size: 12px;
    color: #333;
  }

  .input {
    flex: 1;
    height: 30px;
    border: 1px solid #e2e2dd;
    border-radius: 8px;
    background: #f6f5f1;
    padding: 0 10px;
    font-size: 12px;
    box-sizing: border-box;
    outline: none;
  }

  .picker-wrap {
  flex: 1;              /* 占满右侧 */
  min-width: 0;
}

/* 标题下分割线 */
.card-divider {
  height: 1px;
  background: #CFCFC9;
  margin: 10px 0 12px;
}

/* 每个板块统一外框 */
.block {
  border: 1px solid #E2E2DD;
  background: #FBFBF9;
  border-radius: 10px;
  padding: 10px;
  box-sizing: border-box;
  margin-bottom: 12px;
}

/* 行：label + input 横向排列 */
.row {
  display: flex;
  align-items: center;
  gap: 10px;
}

.row + .row {
  margin-top: 10px; /* 名称/译名两行之间间距 */
}

/* label 左侧固定宽度（你可按视觉调整） */
.label {
  width: 72px;
  font-size: 12px;
  color: #555;
  flex-shrink: 0;
}

/* 输入控件统一 */
.input,
.textarea {
  flex: 1;
  border: 1px solid #E2E2DD;
  background: #F6F5F1;
  border-radius: 8px;
  padding: 8px 10px;
  box-sizing: border-box;
  outline: none;
  font-size: 12px;
  color: #333;
}

/* textarea 行顶对齐 */
.textarea-row {
  align-items: flex-start;
}

.textarea {
  resize: vertical; /* 可选：允许拉伸 */
}

/* 上传封面块内部布局 */
.cover-block {
  display: flex;
  align-items: center;
  gap: 12px;
}

.cover-box {
  width: 120px;
  height: 160px;
  border: 1px dashed #E2E2DD;
  border-radius: 10px;
  background: #F6F5F1;
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
  border: 1px solid #E2E2DD;
  background: #FBFBF9;
  cursor: pointer;
  font-size: 12px;
  color: #333;
  user-select: none;
}

.upload-btn input {
  display: none;
}

  /* 页面内按钮（仅影响本页） */
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
    border: 1px solid #E2E2DD;
  }
  .small {
    height: 26px;
    padding: 0 10px;
    font-size: 12px;
  }
  .icon-btn {
    border: 1px solid #E2E2DD;
    background: #FBFBF9;
    width: 28px;
    height: 28px;
    border-radius: 6px;
    cursor: pointer;
  }

  .page {
    padding: 16px;
    box-sizing: border-box;
  }

  /* 三列横向排列 */
    .grid {
  display: grid;
  grid-template-columns: repeat(3, 460px); /* 等宽 */
  grid-auto-rows: 1fr;                     /* ✅ 关键：等高 */

  gap: 30px;

  justify-content: center; /* 水平居中 */
  align-items: stretch;    /* ✅ 关键：拉伸子项高度 */

  margin: 0 auto;
  padding: 16px;
  box-sizing: border-box;
}



  @media (max-width: 1200px) {
    .grid {
      grid-template-columns: 1fr;
    }
  }

  .card {
    background: #FBFBF9;
    border: 1px solid #E2E2DD;
    border-radius: 10px;
    padding: 14px;
    box-sizing: border-box;
    height: 100%; /* ✅ 必须 */
  }

  .card-title {
    font-size: 13px;
    font-weight: 700;
    color: #333;
  }

  .card-title-row {
    display: flex;
    align-items: center;
    justify-content: space-between;
    margin-bottom: 10px;
  }

  .hint {
    margin-top: 10px;
    font-size: 12px;
    color: #777;
    line-height: 1.4;
  }

  /* 项目信息 */
  .cover-row {
    display: flex;
    gap: 12px;
    align-items: center;
    margin: 12px 0;
  }
  .cover-box {
    width: 120px;
    height: 160px;
    border: 1px dashed #E2E2DD;
    border-radius: 10px;
    background: #F6F5F1;
    display: flex;
    align-items: center;
    justify-content: center;
    overflow: hidden;
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
    border: 1px solid #E2E2DD;
    background: #FBFBF9;
    cursor: pointer;
    font-size: 12px;
    color: #333;
    user-select: none;
  }
  .upload-btn input {
    display: none;
  }

  .form {
    display: grid;
    gap: 10px;
  }
  .field label {
    display: block;
    font-size: 12px;
    color: #555;
    margin-bottom: 6px;
  }
  .field input,
  .field textarea,
  .perm-name,
  .perm-members,
  select {
    width: 100%;
    border: 1px solid #E2E2DD;
    background: #F6F5F1;
    border-radius: 8px;
    padding: 8px 10px;
    box-sizing: border-box;
    outline: none;
    font-size: 12px;
    color: #333;
  }
  .chips {
    display: flex;
    flex-wrap: wrap;
    gap: 8px;
  }
  .chip {
    border: 1px solid #E2E2DD;
    background: #F6F5F1;
    border-radius: 999px;
    padding: 6px 10px;
    font-size: 12px;
    cursor: pointer;
  }
  .chip.active {
    background: #333;
    color: #fff;
    border-color: #333;
  }

  /* 权限 */
  .perm-list {
    display: grid;
    gap: 10px;
  }
  .perm-card {
    border: 1px solid #E2E2DD;
    border-radius: 10px;
    background: #F6F5F1;
    padding: 10px;
  }
  .perm-head {
    display: flex;
    align-items: center;
    gap: 8px;
    margin-bottom: 10px;
  }
  .perm-name {
    flex: 1;
    padding: 8px 10px;
    background: #FBFBF9;
  }
  .perm-row {
    display: grid;
    grid-template-columns: 50px 1fr;
    gap: 10px;
    align-items: center;
    margin-top: 8px;
  }
  .perm-label {
    font-size: 12px;
    color: #666;
  }
  .perm-checks {
    display: flex;
    flex-wrap: wrap;
    gap: 12px;
    font-size: 12px;
    color: #333;
  }

  /* 章节 */
  .chapter-list {
    display: grid;
    gap: 10px;
  }
  .chapter-card {
    border: 1px solid #E2E2DD;
    border-radius: 10px;
    background: #F6F5F1;
    padding: 10px;
  }
  .chapter-head {
    display: flex;
    align-items: flex-start;
    justify-content: space-between;
    gap: 10px;
    margin-bottom: 10px;
  }
  .chapter-left {
    display: flex;
    flex-wrap: wrap;
    gap: 10px;
    align-items: center;
  }
  .chapter-field label {
    display: block;
    font-size: 12px;
    color: #666;
    margin-bottom: 6px;
  }
  .chapter-field input,
  .chapter-field select {
    width: 120px;
    background: #FBFBF9;
  }
  .free-toggle {
    display: inline-flex;
    align-items: center;
    gap: 6px;
    font-size: 12px;
    color: #333;
    margin-top: 18px;
  }

  .price-grid {
    display: grid;
    grid-template-columns: repeat(4, minmax(0, 1fr));
    gap: 8px;
  }
  .price-grid.disabled {
    opacity: 0.55;
  }
  .price-item {
    border: 1px solid #E2E2DD;
    background: #FBFBF9;
    border-radius: 8px;
    padding: 8px;
    display: grid;
    gap: 6px;
  }
  .price-tier {
    font-size: 12px;
    color: #666;
  }
  .price-item input {
    background: #F6F5F1;
  }
</style>
