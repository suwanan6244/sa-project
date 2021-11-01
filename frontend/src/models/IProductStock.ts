import { StaffsInterface } from "./IStaff";
import { SuppliersInterface } from "./ISupplier";
import { ProductsInterface } from "./IProduct";

export interface ProductStockInterface {
  ID: number,
  ProductID: number,
  Product: ProductsInterface,
  Price: number,
  Amount: number,
  StaffID: number,
  Staff: StaffsInterface,
  SupplierID: number,
  Supplier: SuppliersInterface,
  ProductTime: Date,
}