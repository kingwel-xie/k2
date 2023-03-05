export interface SendMessageType {
  checked?: boolean;
  targetType: string;
  targets?: (string | number)[];
  title: string;
  content: string;
}
