// @generated by protoc-gen-connect-es v1.4.0 with parameter "target=ts"
// @generated from file ssoready/v1/ssoready.proto (package ssoready.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { CreateSAMLConnectionRequest, Environment, GetEnvironmentRequest, GetOrganizationRequest, GetSAMLConnectionRequest, GetSAMLRedirectURLRequest, GetSAMLRedirectURLResponse, ListEnvironmentsRequest, ListEnvironmentsResponse, ListOrganizationsRequest, ListOrganizationsResponse, ListSAMLConnectionsRequest, ListSAMLConnectionsResponse, ListSAMLFlowsRequest, ListSAMLFlowsResponse, ListSAMLFlowStepsRequest, ListSAMLFlowStepsResponse, Organization, RedeemSAMLAccessCodeRequest, RedeemSAMLAccessCodeResponse, SAMLConnection, SignInRequest, SignInResponse, UpdateSAMLConnectionRequest, WhoamiRequest, WhoamiResponse } from "./ssoready_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service ssoready.v1.SSOReadyService
 */
export const SSOReadyService = {
  typeName: "ssoready.v1.SSOReadyService",
  methods: {
    /**
     * @generated from rpc ssoready.v1.SSOReadyService.GetSAMLRedirectURL
     */
    getSAMLRedirectURL: {
      name: "GetSAMLRedirectURL",
      I: GetSAMLRedirectURLRequest,
      O: GetSAMLRedirectURLResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc ssoready.v1.SSOReadyService.RedeemSAMLAccessCode
     */
    redeemSAMLAccessCode: {
      name: "RedeemSAMLAccessCode",
      I: RedeemSAMLAccessCodeRequest,
      O: RedeemSAMLAccessCodeResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc ssoready.v1.SSOReadyService.SignIn
     */
    signIn: {
      name: "SignIn",
      I: SignInRequest,
      O: SignInResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc ssoready.v1.SSOReadyService.Whoami
     */
    whoami: {
      name: "Whoami",
      I: WhoamiRequest,
      O: WhoamiResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc ssoready.v1.SSOReadyService.ListEnvironments
     */
    listEnvironments: {
      name: "ListEnvironments",
      I: ListEnvironmentsRequest,
      O: ListEnvironmentsResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc ssoready.v1.SSOReadyService.GetEnvironment
     */
    getEnvironment: {
      name: "GetEnvironment",
      I: GetEnvironmentRequest,
      O: Environment,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc ssoready.v1.SSOReadyService.ListOrganizations
     */
    listOrganizations: {
      name: "ListOrganizations",
      I: ListOrganizationsRequest,
      O: ListOrganizationsResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc ssoready.v1.SSOReadyService.GetOrganization
     */
    getOrganization: {
      name: "GetOrganization",
      I: GetOrganizationRequest,
      O: Organization,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc ssoready.v1.SSOReadyService.ListSAMLConnections
     */
    listSAMLConnections: {
      name: "ListSAMLConnections",
      I: ListSAMLConnectionsRequest,
      O: ListSAMLConnectionsResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc ssoready.v1.SSOReadyService.GetSAMLConnection
     */
    getSAMLConnection: {
      name: "GetSAMLConnection",
      I: GetSAMLConnectionRequest,
      O: SAMLConnection,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc ssoready.v1.SSOReadyService.CreateSAMLConnection
     */
    createSAMLConnection: {
      name: "CreateSAMLConnection",
      I: CreateSAMLConnectionRequest,
      O: SAMLConnection,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc ssoready.v1.SSOReadyService.UpdateSAMLConnection
     */
    updateSAMLConnection: {
      name: "UpdateSAMLConnection",
      I: UpdateSAMLConnectionRequest,
      O: SAMLConnection,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc ssoready.v1.SSOReadyService.ListSAMLFlows
     */
    listSAMLFlows: {
      name: "ListSAMLFlows",
      I: ListSAMLFlowsRequest,
      O: ListSAMLFlowsResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc ssoready.v1.SSOReadyService.ListSAMLFlowSteps
     */
    listSAMLFlowSteps: {
      name: "ListSAMLFlowSteps",
      I: ListSAMLFlowStepsRequest,
      O: ListSAMLFlowStepsResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;

