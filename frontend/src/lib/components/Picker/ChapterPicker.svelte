<script lang="ts">
  import BasePicker, { type PickerItem } from '$lib/components/Picker/BasePicker.svelte';

  export type Chapter = { id: string; name: string };

  let {
    chapters = [] as Chapter[],
    values = [] as string[],
    onChange = (_ids: string[]) => {},
    onSearch
  } = $props<{
    chapters: Chapter[];
    values: string[];
    onChange?: (ids: string[]) => void;
    onSearch?: (q: string) => void | Promise<void>;
  }>();

  // ✅ 让 BasePicker 永远收到“真实数组”
  let selectedSafe = $state<PickerItem[]>([]);

  function recomputeSelected() {
    const chaptersArr: Chapter[] = Array.isArray(chapters) ? chapters : [];
    const valuesArr: string[] = Array.isArray(values) ? values : [];

    const map = new Map<string, Chapter>(chaptersArr.map((c: Chapter) => [c.id, c]));
    selectedSafe = valuesArr
      .map((id: string) => map.get(id))
      .filter((x): x is Chapter => Boolean(x));
  }

  $effect(() => {
    chapters;
    values;
    recomputeSelected();
  });

  const getLabel = (i: PickerItem) => (i as Chapter).name;

  const filterFn = (i: PickerItem, q: string) =>
    (i as Chapter).name.toLowerCase().includes(q);

  function addChapter(i: PickerItem) {
    const id: string = (i as Chapter).id;
    const valuesArr: string[] = Array.isArray(values) ? values : [];
    if (valuesArr.includes(id)) return;
    onChange([...valuesArr, id]);
  }

  function removeChapter(id: string) {
    const valuesArr: string[] = Array.isArray(values) ? values : [];
    onChange(valuesArr.filter((x: string) => x !== id));
  }
</script>

<BasePicker
  items={chapters}
  selected={selectedSafe}
  single={false}
  onAdd={addChapter}
  onRemove={removeChapter}
  getLabel={getLabel}
  filterFn={filterFn}
  onSearch={onSearch}
  placeholder="搜索章节"
/>
