import { UseFormReturnType } from '/@/components/Form';
import { Ref } from 'vue';

export async function validateTabForms(
  active: Ref<string>,
  tabKeys: string[],
  forms: UseFormReturnType[],
): Promise<any[]> {
  return validateTabs(
    active,
    tabKeys,
    forms.map(([, { validate }]) => validate()),
  );
}

type Task = Promise<any> | undefined;

export async function validateTabs(
  active: Ref<string>,
  tabKeys: string[],
  tasks: Task[],
): Promise<any[]> {
  const rets: any[] = [];
  // 并行验证, 遇到失败就结束
  for (let i = 0; i < tasks.length; i++) {
    if (tasks[i]) {
      try {
        rets.push(await tasks[i]);
      } catch (e) {
        active.value = tabKeys[i];
        throw e;
      }
    }
  }
  return rets;
}
