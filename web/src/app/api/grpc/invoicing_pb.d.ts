import * as jspb from 'google-protobuf'

import * as google_protobuf_empty_pb from 'google-protobuf/google/protobuf/empty_pb'; // proto import: "google/protobuf/empty.proto"
import * as common_pb from './common_pb'; // proto import: "common.proto"


export class Invoice extends jspb.Message {
  getId(): string;
  setId(value: string): Invoice;
  hasId(): boolean;
  clearId(): Invoice;

  getClientId(): string;
  setClientId(value: string): Invoice;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Invoice.AsObject;
  static toObject(includeInstance: boolean, msg: Invoice): Invoice.AsObject;
  static serializeBinaryToWriter(message: Invoice, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Invoice;
  static deserializeBinaryFromReader(message: Invoice, reader: jspb.BinaryReader): Invoice;
}

export namespace Invoice {
  export type AsObject = {
    id?: string,
    clientId: string,
  }

  export enum IdCase { 
    _ID_NOT_SET = 0,
    ID = 1,
  }
}

export class ListInvoices extends jspb.Message {
  getInvoicesList(): Array<Invoice>;
  setInvoicesList(value: Array<Invoice>): ListInvoices;
  clearInvoicesList(): ListInvoices;
  addInvoices(value?: Invoice, index?: number): Invoice;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListInvoices.AsObject;
  static toObject(includeInstance: boolean, msg: ListInvoices): ListInvoices.AsObject;
  static serializeBinaryToWriter(message: ListInvoices, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListInvoices;
  static deserializeBinaryFromReader(message: ListInvoices, reader: jspb.BinaryReader): ListInvoices;
}

export namespace ListInvoices {
  export type AsObject = {
    invoicesList: Array<Invoice.AsObject>,
  }
}

