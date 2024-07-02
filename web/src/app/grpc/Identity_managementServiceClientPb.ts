/**
 * @fileoverview gRPC-Web generated client stub for hestia.v1
 * @enhanceable
 * @public
 */

// Code generated by protoc-gen-grpc-web. DO NOT EDIT.
// versions:
// 	protoc-gen-grpc-web v1.5.0
// 	protoc              v5.26.1
// source: identity_management.proto


/* eslint-disable */
// @ts-nocheck


import * as grpcWeb from 'grpc-web';

import * as google_protobuf_empty_pb from 'google-protobuf/google/protobuf/empty_pb'; // proto import: "google/protobuf/empty.proto"
import * as identity_management_pb from './identity_management_pb'; // proto import: "identity_management.proto"


export class IdentityManagementClient {
  client_: grpcWeb.AbstractClientBase;
  hostname_: string;
  credentials_: null | { [index: string]: string; };
  options_: null | { [index: string]: any; };

  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; }) {
    if (!options) options = {};
    if (!credentials) credentials = {};
    options['format'] = 'text';

    this.client_ = new grpcWeb.GrpcWebClientBase(options);
    this.hostname_ = hostname.replace(/\/+$/, '');
    this.credentials_ = credentials;
    this.options_ = options;
  }

  methodDescriptorLogin = new grpcWeb.MethodDescriptor(
    '/hestia.v1.IdentityManagement/Login',
    grpcWeb.MethodType.UNARY,
    identity_management_pb.LoginRequest,
    identity_management_pb.LoginResponse,
    (request: identity_management_pb.LoginRequest) => {
      return request.serializeBinary();
    },
    identity_management_pb.LoginResponse.deserializeBinary
  );

  login(
    request: identity_management_pb.LoginRequest,
    metadata?: grpcWeb.Metadata | null): Promise<identity_management_pb.LoginResponse>;

  login(
    request: identity_management_pb.LoginRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: identity_management_pb.LoginResponse) => void): grpcWeb.ClientReadableStream<identity_management_pb.LoginResponse>;

  login(
    request: identity_management_pb.LoginRequest,
    metadata?: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: identity_management_pb.LoginResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/hestia.v1.IdentityManagement/Login',
        request,
        metadata || {},
        this.methodDescriptorLogin,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/hestia.v1.IdentityManagement/Login',
    request,
    metadata || {},
    this.methodDescriptorLogin);
  }

  methodDescriptorRegister = new grpcWeb.MethodDescriptor(
    '/hestia.v1.IdentityManagement/Register',
    grpcWeb.MethodType.UNARY,
    identity_management_pb.RegisterRequest,
    google_protobuf_empty_pb.Empty,
    (request: identity_management_pb.RegisterRequest) => {
      return request.serializeBinary();
    },
    google_protobuf_empty_pb.Empty.deserializeBinary
  );

  register(
    request: identity_management_pb.RegisterRequest,
    metadata?: grpcWeb.Metadata | null): Promise<google_protobuf_empty_pb.Empty>;

  register(
    request: identity_management_pb.RegisterRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: google_protobuf_empty_pb.Empty) => void): grpcWeb.ClientReadableStream<google_protobuf_empty_pb.Empty>;

  register(
    request: identity_management_pb.RegisterRequest,
    metadata?: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: google_protobuf_empty_pb.Empty) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/hestia.v1.IdentityManagement/Register',
        request,
        metadata || {},
        this.methodDescriptorRegister,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/hestia.v1.IdentityManagement/Register',
    request,
    metadata || {},
    this.methodDescriptorRegister);
  }

  methodDescriptorAlive = new grpcWeb.MethodDescriptor(
    '/hestia.v1.IdentityManagement/Alive',
    grpcWeb.MethodType.UNARY,
    identity_management_pb.TokenRequest,
    google_protobuf_empty_pb.Empty,
    (request: identity_management_pb.TokenRequest) => {
      return request.serializeBinary();
    },
    google_protobuf_empty_pb.Empty.deserializeBinary
  );

  alive(
    request: identity_management_pb.TokenRequest,
    metadata?: grpcWeb.Metadata | null): Promise<google_protobuf_empty_pb.Empty>;

  alive(
    request: identity_management_pb.TokenRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: google_protobuf_empty_pb.Empty) => void): grpcWeb.ClientReadableStream<google_protobuf_empty_pb.Empty>;

  alive(
    request: identity_management_pb.TokenRequest,
    metadata?: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: google_protobuf_empty_pb.Empty) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/hestia.v1.IdentityManagement/Alive',
        request,
        metadata || {},
        this.methodDescriptorAlive,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/hestia.v1.IdentityManagement/Alive',
    request,
    metadata || {},
    this.methodDescriptorAlive);
  }

  methodDescriptorLogout = new grpcWeb.MethodDescriptor(
    '/hestia.v1.IdentityManagement/Logout',
    grpcWeb.MethodType.UNARY,
    identity_management_pb.TokenRequest,
    google_protobuf_empty_pb.Empty,
    (request: identity_management_pb.TokenRequest) => {
      return request.serializeBinary();
    },
    google_protobuf_empty_pb.Empty.deserializeBinary
  );

  logout(
    request: identity_management_pb.TokenRequest,
    metadata?: grpcWeb.Metadata | null): Promise<google_protobuf_empty_pb.Empty>;

  logout(
    request: identity_management_pb.TokenRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: google_protobuf_empty_pb.Empty) => void): grpcWeb.ClientReadableStream<google_protobuf_empty_pb.Empty>;

  logout(
    request: identity_management_pb.TokenRequest,
    metadata?: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: google_protobuf_empty_pb.Empty) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/hestia.v1.IdentityManagement/Logout',
        request,
        metadata || {},
        this.methodDescriptorLogout,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/hestia.v1.IdentityManagement/Logout',
    request,
    metadata || {},
    this.methodDescriptorLogout);
  }

}

