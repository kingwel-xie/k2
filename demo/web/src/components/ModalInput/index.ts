import { withInstall } from '/@/utils';
import modalInput from './src/ModalInput.vue';

export const ModalInput = withInstall(modalInput);

export type InputCallback = (val: Recordable) => void;

export interface PromptModalAction {
  openModal: (ok?: InputCallback, cancel?: InputCallback) => void;
}
