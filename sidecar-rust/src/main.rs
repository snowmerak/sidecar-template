// use tonic::transport::Server;
// use core_proto::company_auth_v1::auth_service_server::AuthServiceServer;

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    // let addr = "0.0.0.0:50052".parse()?;
    
    // println!("Sidecar Rust Server listening on {}", addr);

    // Server::builder()
        // TODO: Add your services here
        // .add_service(AuthServiceServer::new(MyAuthService::default()))
        // .serve(addr)
        // .await?;

    Ok(())
}
