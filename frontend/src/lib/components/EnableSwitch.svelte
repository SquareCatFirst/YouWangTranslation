<script lang="ts">
  let {
    value = true, // true=启用（圆在左），false=禁用（圆在右）
    onLabel = '启用',
    offLabel = '禁用',
    disabled = false,
    onChange = (_v: boolean) => {}
  } = $props<{
    value: boolean;
    onLabel?: string;
    offLabel?: string;
    disabled?: boolean;
    onChange?: (v: boolean) => void;
  }>();

  function toggle() {
    if (disabled) return;
    onChange(!value);
  }

  function onKeydown(e: KeyboardEvent) {
    if (disabled) return;
    if (e.key === 'Enter' || e.key === ' ') {
      e.preventDefault();
      toggle();
    }
  }
</script>

<button
  type="button"
  class="switch"
  class:on={value}
  class:off={!value}
  disabled={disabled}
  role="switch"
  aria-checked={value}
  onclick={toggle}
  onkeydown={onKeydown}
>
  <!-- 左文字区（禁用时显示） -->
  <span class="text left">{offLabel}</span>

  <!-- 圆点 -->
  <span class="knob" aria-hidden="true"></span>

  <!-- 右文字区（启用时显示） -->
  <span class="text right">{onLabel}</span>
</button>

<style>
  .switch {
  height: 30px;
  min-width: 92px;

  border: 1px solid #797979;
  border-radius: 999px;

  display: inline-flex;
  align-items: center;

  padding: 0px;
  box-sizing: border-box;

  cursor: pointer;
  user-select: none;

  background: #7f8f86;
  position: relative;
  overflow: hidden;

  transition: background-color 180ms ease;

  /* ✅ 关键：给内容区留出左右 padding，圆点用 absolute */
  padding: 0 6px;
}

/* 文字保留占位 */
.text {
  font-size: 12px;
  color: #ffffff;
  white-space: nowrap;
  line-height: 1;
  width: 34px;
  text-align: center;
  transition: opacity 160ms ease, transform 160ms ease;
}

/* ✅ 关键：圆点 absolute，靠 transform 移动 */
.knob {
  position: absolute;
  top: 0;
  left: 0;

  width: 32px;
  height: 30px;
  border-radius: 999px;
  background: #ffffff;
  border: 1px solid #797979;
  box-sizing: border-box;

  transition: transform 220ms cubic-bezier(0.2, 0.8, 0.2, 1);
  will-change: transform;

  /* ✅ 默认在左（on） */
  transform: translate(-1px,-1px);
}

/* ===== 状态：启用（圆在左） ===== */
.switch.on {
  background: #7f8f86;
}

/* ===== 状态：禁用（圆在右） ===== */
.switch.off {
  background: #c27a7a;
}

/* ✅ 关键：把圆点移动到最右边
   92px(min-width) - 32px(knob) = 60px
   再减去边框等细节，你可按视觉微调 60/58/62
*/
.switch.off .knob {
  transform: translate(59px,-1px);
}

/* 文字淡入淡出（你原本的逻辑可以继续用） */
.switch.on .left {
  opacity: 0;
  transform: translateX(-6px);
  pointer-events: none;
}
.switch.on .right {
  opacity: 1;
  transform: translateX(0);
}

.switch.off .right {
  opacity: 0;
  transform: translateX(6px);
  pointer-events: none;
}
.switch.off .left {
  opacity: 1;
  transform: translateX(0);
}

</style>
