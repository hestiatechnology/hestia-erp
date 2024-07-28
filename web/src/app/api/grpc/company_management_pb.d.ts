import * as jspb from 'google-protobuf'

// eslint-disable-next-line @typescript-eslint/no-unused-vars
import * as google_protobuf_empty_pb from 'google-protobuf/google/protobuf/empty_pb'; // proto import: "google/protobuf/empty.proto"


export class CreateCompanyRequest extends jspb.Message {
  getName(): string;
  setName(value: string): CreateCompanyRequest;

  getIssoletrader(): boolean;
  setIssoletrader(value: boolean): CreateCompanyRequest;

  getCommercialname(): string;
  setCommercialname(value: string): CreateCompanyRequest;
  hasCommercialname(): boolean;
  clearCommercialname(): CreateCompanyRequest;

  getVatid(): number;
  setVatid(value: number): CreateCompanyRequest;

  getSsn(): number;
  setSsn(value: number): CreateCompanyRequest;

  getLocation(): Location | undefined;
  setLocation(value?: Location): CreateCompanyRequest;
  hasLocation(): boolean;
  clearLocation(): CreateCompanyRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateCompanyRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CreateCompanyRequest): CreateCompanyRequest.AsObject;
  static serializeBinaryToWriter(message: CreateCompanyRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateCompanyRequest;
  static deserializeBinaryFromReader(message: CreateCompanyRequest, reader: jspb.BinaryReader): CreateCompanyRequest;
}

export namespace CreateCompanyRequest {
  export type AsObject = {
    name: string,
    issoletrader: boolean,
    commercialname?: string,
    vatid: number,
    ssn: number,
    location?: Location.AsObject,
  }

  export enum CommercialnameCase {
    _COMMERCIALNAME_NOT_SET = 0,
    COMMERCIALNAME = 3,
  }
}

export class Location extends jspb.Message {
  getAddress(): string;
  setAddress(value: string): Location;

  getLocality(): string;
  setLocality(value: string): Location;

  getPostalcode(): string;
  setPostalcode(value: string): Location;

  getCountry(): string;
  setCountry(value: string): Location;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Location.AsObject;
  static toObject(includeInstance: boolean, msg: Location): Location.AsObject;
  static serializeBinaryToWriter(message: Location, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Location;
  static deserializeBinaryFromReader(message: Location, reader: jspb.BinaryReader): Location;
}

export namespace Location {
  export type AsObject = {
    address: string,
    locality: string,
    postalcode: string,
    country: string,
  }
}

export class Id extends jspb.Message {
  getId(): string;
  setId(value: string): Id;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Id.AsObject;
  static toObject(includeInstance: boolean, msg: Id): Id.AsObject;
  static serializeBinaryToWriter(message: Id, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Id;
  static deserializeBinaryFromReader(message: Id, reader: jspb.BinaryReader): Id;
}

export namespace Id {
  export type AsObject = {
    id: string,
  }
}

export class UpdateCompanyRequest extends jspb.Message {
  getId(): string;
  setId(value: string): UpdateCompanyRequest;

  getName(): string;
  setName(value: string): UpdateCompanyRequest;

  getDescription(): string;
  setDescription(value: string): UpdateCompanyRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateCompanyRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateCompanyRequest): UpdateCompanyRequest.AsObject;
  static serializeBinaryToWriter(message: UpdateCompanyRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateCompanyRequest;
  static deserializeBinaryFromReader(message: UpdateCompanyRequest, reader: jspb.BinaryReader): UpdateCompanyRequest;
}

export namespace UpdateCompanyRequest {
  export type AsObject = {
    id: string,
    name: string,
    description: string,
  }
}

export class UpdateCompanyResponse extends jspb.Message {
  getId(): string;
  setId(value: string): UpdateCompanyResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateCompanyResponse.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateCompanyResponse): UpdateCompanyResponse.AsObject;
  static serializeBinaryToWriter(message: UpdateCompanyResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateCompanyResponse;
  static deserializeBinaryFromReader(message: UpdateCompanyResponse, reader: jspb.BinaryReader): UpdateCompanyResponse;
}

export namespace UpdateCompanyResponse {
  export type AsObject = {
    id: string,
  }
}

export class ListCompaniesResponse extends jspb.Message {
  getCompaniesList(): Array<Company>;
  setCompaniesList(value: Array<Company>): ListCompaniesResponse;
  clearCompaniesList(): ListCompaniesResponse;
  addCompanies(value?: Company, index?: number): Company;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListCompaniesResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ListCompaniesResponse): ListCompaniesResponse.AsObject;
  static serializeBinaryToWriter(message: ListCompaniesResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListCompaniesResponse;
  static deserializeBinaryFromReader(message: ListCompaniesResponse, reader: jspb.BinaryReader): ListCompaniesResponse;
}

export namespace ListCompaniesResponse {
  export type AsObject = {
    companiesList: Array<Company.AsObject>,
  }
}

export class Company extends jspb.Message {
  getId(): string;
  setId(value: string): Company;

  getName(): string;
  setName(value: string): Company;

  getCommercialname(): string;
  setCommercialname(value: string): Company;

  getDescription(): string;
  setDescription(value: string): Company;

  getLogo(): string;
  setLogo(value: string): Company;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Company.AsObject;
  static toObject(includeInstance: boolean, msg: Company): Company.AsObject;
  static serializeBinaryToWriter(message: Company, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Company;
  static deserializeBinaryFromReader(message: Company, reader: jspb.BinaryReader): Company;
}

export namespace Company {
  export type AsObject = {
    id: string,
    name: string,
    commercialname: string,
    description: string,
    logo: string,
  }
}

export class AddUserToCompanyRequest extends jspb.Message {
  getEmail(): string;
  setEmail(value: string): AddUserToCompanyRequest;

  getCompanyid(): string;
  setCompanyid(value: string): AddUserToCompanyRequest;

  getEmployeeid(): string;
  setEmployeeid(value: string): AddUserToCompanyRequest;
  hasEmployeeid(): boolean;
  clearEmployeeid(): AddUserToCompanyRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AddUserToCompanyRequest.AsObject;
  static toObject(includeInstance: boolean, msg: AddUserToCompanyRequest): AddUserToCompanyRequest.AsObject;
  static serializeBinaryToWriter(message: AddUserToCompanyRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AddUserToCompanyRequest;
  static deserializeBinaryFromReader(message: AddUserToCompanyRequest, reader: jspb.BinaryReader): AddUserToCompanyRequest;
}

export namespace AddUserToCompanyRequest {
  export type AsObject = {
    email: string,
    companyid: string,
    employeeid?: string,
  }

  export enum EmployeeidCase {
    _EMPLOYEEID_NOT_SET = 0,
    EMPLOYEEID = 3,
  }
}

