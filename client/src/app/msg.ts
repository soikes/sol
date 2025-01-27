export enum MsgType {
  Input = 0,
  Transform,
  Spawn,
  Register,
  Sync,
}

export type Msg = {
  type: MsgType;
  seq: Number;
  ts: Number;
  data: any;
};
