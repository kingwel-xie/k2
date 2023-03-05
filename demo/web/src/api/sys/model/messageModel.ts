export interface MessageListItem {
  id: number;
  type: string;
  sender: string;
  receiver: string;
  originId?: number;
  title: string;
  content: string;
  read?: boolean;
  createdAt: string;
}

export interface GetUnreadInfo {
  numNotices: number;
  noticeList: MessageListItem[];
  numMessages: number;
  messageList: MessageListItem[];
}
