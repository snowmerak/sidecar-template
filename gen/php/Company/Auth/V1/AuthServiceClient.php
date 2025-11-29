<?php
// GENERATED CODE -- DO NOT EDIT!

namespace Company\Auth\V1;

/**
 */
class AuthServiceClient extends \Grpc\BaseStub {

    /**
     * @param string $hostname hostname
     * @param array $opts channel options
     * @param \Grpc\Channel $channel (optional) re-use channel object
     */
    public function __construct($hostname, $opts, $channel = null) {
        parent::__construct($hostname, $opts, $channel);
    }

    /**
     * @param \Company\Auth\V1\LoginRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall<\Company\Auth\V1\LoginResponse>
     */
    public function Login(\Company\Auth\V1\LoginRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/company.auth.v1.AuthService/Login',
        $argument,
        ['\Company\Auth\V1\LoginResponse', 'decode'],
        $metadata, $options);
    }

}
