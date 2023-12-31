# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from . import OpenaiActions_pb2

class SchoolServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.GetKey = channel.unary_unary(
                '/proto.SchoolService/GetKey',
                request_serializer=OpenaiActions_pb2.KeyRequest.SerializeToString,
                response_deserializer=OpenaiActions_pb2.KeyResponse.FromString,
                )


class SchoolServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def GetKey(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_SchoolServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'GetKey': grpc.unary_unary_rpc_method_handler(
                    servicer.GetKey,
                    request_deserializer=OpenaiActions__pb2.KeyRequest.FromString,
                    response_serializer=OpenaiActions__pb2.KeyResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'proto.SchoolService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class SchoolService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def GetKey(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/proto.SchoolService/GetKey',
            OpenaiActions__pb2.KeyRequest.SerializeToString,
            OpenaiActions__pb2.KeyResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
