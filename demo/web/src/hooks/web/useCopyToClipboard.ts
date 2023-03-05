import { computed, ComputedRef, Ref, ref, watch } from 'vue';
import { MaybeComputedRef, resolveUnref, tryOnMounted, useTimeoutFn } from '@vueuse/shared';
import { isDef } from '/@/utils/is';
interface Options {
  target?: HTMLElement;
}
export function useCopyToClipboard(initial?: string) {
  const clipboardRef = ref(initial || '');
  const isSuccessRef = ref(false);
  const copiedRef = ref(false);

  watch(
    clipboardRef,
    (str?: string) => {
      if (isDef(str)) {
        copiedRef.value = true;
        isSuccessRef.value = copyTextToClipboard(str);
      }
    },
    { immediate: !!initial, flush: 'sync' },
  );

  return { clipboardRef, isSuccessRef, copiedRef };
}

export function useSupported(callback: () => unknown, sync = false) {
  const isSupported = ref() as Ref<boolean>;

  const update = () => (isSupported.value = Boolean(callback()));

  update();

  tryOnMounted(update, sync);
  return isSupported;
}

export interface UseClipboardOptions<Source> {
  /**
   * Copy source
   */
  source?: Source;

  /**
   * Milliseconds to reset state of `copied` ref
   *
   * @default 1500
   */
  copiedDuring?: number;

  /**
   * Whether fallback to document.execCommand('copy') if clipboard is undefined.
   *
   * @default false
   */
  legacy?: boolean;

  /**
   * call after copied
   * @param text
   */
  onCopied?: (text: string) => void;

  /**
   * call after fail
   * @param text
   */
  onFail?: (e) => void;
}

export interface UseClipboardReturn<Optional> {
  isSupported: Ref<boolean>;
  text: ComputedRef<string>;
  copied: ComputedRef<boolean>;
  copy: Optional extends true ? (text?: string) => Promise<void> : (text: string) => Promise<void>;
}

/**
 * Reactive Clipboard API.
 *
 * @see https://vueuse.org/useClipboard
 * @param options
 */
export function useClipboard(options?: UseClipboardOptions<undefined>): UseClipboardReturn<false>;
export function useClipboard(
  options: UseClipboardOptions<MaybeComputedRef<string>>,
): UseClipboardReturn<true>;
export function useClipboard(
  options: UseClipboardOptions<MaybeComputedRef<string> | undefined> = {},
): UseClipboardReturn<boolean> {
  const { source, copiedDuring = 1500, legacy = true, onCopied, onFail } = options;
  const isClipboardApiSupported = useSupported(() => navigator && 'clipboard' in navigator);
  const isSupported = computed(() => isClipboardApiSupported.value || legacy);
  const text = ref('');
  const copied = ref(false);
  const timeout = useTimeoutFn(() => (copied.value = false), copiedDuring);

  async function copy(value = resolveUnref(source)) {
    if (isSupported.value && value != null) {
      try {
        if (isClipboardApiSupported.value) {
          await navigator!.clipboard.writeText(value);
        } else {
          copyTextToClipboard(value);
        }
      } catch (e) {
        onFail?.(e);
        console.error(e);
      }
      text.value = value;
      copied.value = true;
      onCopied?.(value);
      timeout.start();
    }
  }

  return {
    isSupported,
    text: text as ComputedRef<string>,
    copied: copied as ComputedRef<boolean>,
    copy,
  };
}

export function copyTextToClipboard(input: string, { target = document.body }: Options = {}) {
  const element = document.createElement('textarea');
  const previouslyFocusedElement = document.activeElement;

  element.value = input;

  element.setAttribute('readonly', '');

  (element.style as any).contain = 'strict';
  element.style.position = 'absolute';
  element.style.left = '-9999px';
  element.style.fontSize = '12pt';

  const selection = document.getSelection();
  let originalRange;
  if (selection && selection.rangeCount > 0) {
    originalRange = selection.getRangeAt(0);
  }

  target.append(element);
  element.select();

  element.selectionStart = 0;
  element.selectionEnd = input.length;

  let isSuccess = false;
  try {
    isSuccess = document.execCommand('copy');
  } catch (e: any) {
    throw new Error(e);
  }

  element.remove();

  if (originalRange && selection) {
    selection.removeAllRanges();
    selection.addRange(originalRange);
  }

  if (previouslyFocusedElement) {
    (previouslyFocusedElement as HTMLElement).focus();
  }
  return isSuccess;
}
