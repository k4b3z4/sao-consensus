/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "saonetwork.sao.sao";

export interface Shard {
  id: number;
  orderId: number;
  status: number;
  size: number;
  cid: string;
  pledge: number;
}

const baseShard: object = {
  id: 0,
  orderId: 0,
  status: 0,
  size: 0,
  cid: "",
  pledge: 0,
};

export const Shard = {
  encode(message: Shard, writer: Writer = Writer.create()): Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    if (message.orderId !== 0) {
      writer.uint32(16).uint64(message.orderId);
    }
    if (message.status !== 0) {
      writer.uint32(24).int32(message.status);
    }
    if (message.size !== 0) {
      writer.uint32(32).int32(message.size);
    }
    if (message.cid !== "") {
      writer.uint32(42).string(message.cid);
    }
    if (message.pledge !== 0) {
      writer.uint32(48).uint64(message.pledge);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Shard {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseShard } as Shard;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.orderId = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.status = reader.int32();
          break;
        case 4:
          message.size = reader.int32();
          break;
        case 5:
          message.cid = reader.string();
          break;
        case 6:
          message.pledge = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Shard {
    const message = { ...baseShard } as Shard;
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    if (object.orderId !== undefined && object.orderId !== null) {
      message.orderId = Number(object.orderId);
    } else {
      message.orderId = 0;
    }
    if (object.status !== undefined && object.status !== null) {
      message.status = Number(object.status);
    } else {
      message.status = 0;
    }
    if (object.size !== undefined && object.size !== null) {
      message.size = Number(object.size);
    } else {
      message.size = 0;
    }
    if (object.cid !== undefined && object.cid !== null) {
      message.cid = String(object.cid);
    } else {
      message.cid = "";
    }
    if (object.pledge !== undefined && object.pledge !== null) {
      message.pledge = Number(object.pledge);
    } else {
      message.pledge = 0;
    }
    return message;
  },

  toJSON(message: Shard): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.orderId !== undefined && (obj.orderId = message.orderId);
    message.status !== undefined && (obj.status = message.status);
    message.size !== undefined && (obj.size = message.size);
    message.cid !== undefined && (obj.cid = message.cid);
    message.pledge !== undefined && (obj.pledge = message.pledge);
    return obj;
  },

  fromPartial(object: DeepPartial<Shard>): Shard {
    const message = { ...baseShard } as Shard;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    if (object.orderId !== undefined && object.orderId !== null) {
      message.orderId = object.orderId;
    } else {
      message.orderId = 0;
    }
    if (object.status !== undefined && object.status !== null) {
      message.status = object.status;
    } else {
      message.status = 0;
    }
    if (object.size !== undefined && object.size !== null) {
      message.size = object.size;
    } else {
      message.size = 0;
    }
    if (object.cid !== undefined && object.cid !== null) {
      message.cid = object.cid;
    } else {
      message.cid = "";
    }
    if (object.pledge !== undefined && object.pledge !== null) {
      message.pledge = object.pledge;
    } else {
      message.pledge = 0;
    }
    return message;
  },
};

declare var self: any | undefined;
declare var window: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") return globalThis;
  if (typeof self !== "undefined") return self;
  if (typeof window !== "undefined") return window;
  if (typeof global !== "undefined") return global;
  throw "Unable to locate global object";
})();

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (util.Long !== Long) {
  util.Long = Long as any;
  configure();
}