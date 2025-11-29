// @generated
// This file is manually maintained to export generated modules.
// When you add a new proto package, add a corresponding module here.

#[path = "company.auth.v1.rs"]
pub mod company_auth_v1;

// Re-export for convenience if needed
pub mod company {
    pub mod auth {
        pub mod v1 {
            pub use crate::company_auth_v1::*;
        }
    }
}
