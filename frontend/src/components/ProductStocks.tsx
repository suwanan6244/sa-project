import { useEffect, useState } from "react";
import { Link as RouterLink } from "react-router-dom";
import { createStyles, makeStyles, Theme } from "@material-ui/core/styles";
import Typography from "@material-ui/core/Typography";
import Button from "@material-ui/core/Button";
import Container from "@material-ui/core/Container";
import Paper from "@material-ui/core/Paper";
import Box from "@material-ui/core/Box";
import Table from "@material-ui/core/Table";
import TableBody from "@material-ui/core/TableBody";
import TableCell from "@material-ui/core/TableCell";
import TableContainer from "@material-ui/core/TableContainer";
import TableHead from "@material-ui/core/TableHead";
import TableRow from "@material-ui/core/TableRow";
import { ProductStockInterface } from "../models/IProductStock";
import { format } from 'date-fns'

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    container: {
      marginTop: theme.spacing(2),
    },
    table: {
      minWidth: 650,
    },
    tableSpace: {
      marginTop: 20,
    },
  })
);

function ProductStocks() {
  const classes = useStyles();
  const [productStocks, setProductStocks] = useState<ProductStockInterface[]>([]);
  const apiUrl = "http://localhost:8080";
  const requestOptions = {
    method: "GET",
    headers: {
      Authorization: `Bearer ${localStorage.getItem("token")}`,
      "Content-Type": "application/json",
    },
  };

  const getProductStocks = async () => {
    fetch(`${apiUrl}/product_stocks`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        console.log(res.data);
        if (res.data) {
          setProductStocks(res.data);
        } else {
          console.log("else");
        }
      });
  };

  useEffect(() => {
    getProductStocks();
  }, []);

  return (
    <div>
      <Container className={classes.container} maxWidth="lg">
        <Box display="flex">
          <Box flexGrow={1}>
            <Typography
              component="h2"
              variant="h6"
              color="primary"
              gutterBottom
            >
              ข้อมูลการบันทึกรายการสินค้า
            </Typography>
          </Box>
          <Box>
            <Button
              component={RouterLink}
              to="/product_stock/create"
              variant="contained"
              color="primary"
            >
              สร้างข้อมูล
            </Button>
          </Box>
        </Box>
        <TableContainer component={Paper} className={classes.tableSpace}>
          <Table className={classes.table} aria-label="simple table">
            <TableHead>
              <TableRow>
                <TableCell align="center" width="10%">
                  ลำดับ
                </TableCell>
                <TableCell align="center" width="15%">
                  ชื่อสินค้า
                </TableCell>
                <TableCell align="center" width="5%">
                  ราคา
                </TableCell>
                <TableCell align="center" width="5%">
                  จำนวน
                </TableCell>
                <TableCell align="center" width="5%">
                  ผู้ผลิต
                </TableCell>
                <TableCell align="center" width="5%">
                  พนักงานเพิ่มสินค้า
                </TableCell>
                <TableCell align="center" width="15%">
                  วันที่และเวลา
                </TableCell>
              </TableRow>
            </TableHead>
            <TableBody>
              {productStocks.map((item: ProductStockInterface) => (
                <TableRow key={item.ID}>
                  <TableCell align="center">{item.ID}</TableCell>
                  <TableCell align="center">{item.Product.Name}</TableCell>
                  <TableCell align="center">{item.Price}</TableCell>
                  <TableCell align="center">{item.Amount}</TableCell>
                  <TableCell align="center">{item.Supplier.Name}</TableCell>
                  <TableCell align="center">{item.Staff.Name}</TableCell>
                  <TableCell align="center">{format((new Date(item.ProductTime)), 'dd MMMM yyyy hh:mm a')}
                  </TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
        </TableContainer>
      </Container>
    </div>
  );
}
export default ProductStocks;