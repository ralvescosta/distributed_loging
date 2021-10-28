use domain::entities::product_entity::ProductEntity;
use serde::{Deserialize, Serialize};

#[derive(Serialize, Deserialize)]
pub struct ProductDocument {
    pub id: String,
    pub product_category: String,
    pub tag: String,
    pub title: String,
    pub subtitle: String,
    pub authors: Vec<String>,
    pub amount_in_stock: i32,
    pub created_at: String,
    pub updated_at: String,
    pub num_pages: i32,
    pub tags: Vec<String>,
}

impl ProductDocument {
    pub fn to_entity(self) -> ProductEntity {
        ProductEntity {
            id: self.id,
            product_category: self.product_category,
            tag: self.tag,
            title: self.title,
            subtitle: self.subtitle,
            authors: self.authors,
            amount_in_stock: self.amount_in_stock,
            created_at: self.created_at,
            updated_at: self.updated_at,
            num_pages: self.num_pages,
            tags: self.tags,
        }
    }
}
