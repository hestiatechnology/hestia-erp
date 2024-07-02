import * as jspb from 'google-protobuf'



export class TechnicalModel extends jspb.Message {
  getId(): string;
  setId(value: string): TechnicalModel;
  hasId(): boolean;
  clearId(): TechnicalModel;

  getName(): string;
  setName(value: string): TechnicalModel;

  getDescription(): string;
  setDescription(value: string): TechnicalModel;

  getType(): string;
  setType(value: string): TechnicalModel;

  getCreated(): string;
  setCreated(value: string): TechnicalModel;

  getUpdated(): string;
  setUpdated(value: string): TechnicalModel;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): TechnicalModel.AsObject;
  static toObject(includeInstance: boolean, msg: TechnicalModel): TechnicalModel.AsObject;
  static serializeBinaryToWriter(message: TechnicalModel, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): TechnicalModel;
  static deserializeBinaryFromReader(message: TechnicalModel, reader: jspb.BinaryReader): TechnicalModel;
}

export namespace TechnicalModel {
  export type AsObject = {
    id?: string,
    name: string,
    description: string,
    type: string,
    created: string,
    updated: string,
  }

  export enum IdCase { 
    _ID_NOT_SET = 0,
    ID = 1,
  }
}

