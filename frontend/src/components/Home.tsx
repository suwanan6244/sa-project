import { createStyles, makeStyles, Theme } from "@material-ui/core/styles";
import Container from "@material-ui/core/Container";

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

function Home() {
  const classes = useStyles();

  return (
    <div>
      <Container className={classes.container} maxWidth="md">
        <h1 style={{ textAlign: "center" }}>ระบบบันทึกการเข้าชมวีดีโอ</h1>
        <h4>Requirements</h4>
        <p>
        ระบบ Farm Mart เป็นระบบที่ไว้ใช้ซื้อสินค้าออนไลน์ซึ่งเป็นผลิตผลทางการเกษตร
        โดยสมาชิกที่เป็นผู้ใช้ระบบสามารถ Login เข้ามาเพื่อเลือกซื้อสินค้าและสั่งซื้อสินค้าที่ต้องการได้ 
        สมาชิกระบบจะสามารถตรวจสอบดูได้ว่ามีสินค้าชนิดไหนที่วางขายอยู่บ้าง 
        โดยจะมีระบบคลังสินค้าไว้เก็บข้อมูลของสินค้าประกอบด้วย รหัสสินค้า ราคา 
        จำนวนสินค้าที่มีอยู่ในคลัง เป็นต้น ในระบบคลังสินค้าพนักงานจะสามารถเพิ่ม 
        รายการสินค้าและสามารถเลือกประเภทของสินค้าได้ 
        เมื่อพนักงานทำการเพิ่มข้อมูลสินค้าลงไประบบจะทำการบันทึกข้อมูลรายการสินค้า
        หลังจากนั้นจะแสดงข้อความขึ้นมาว่า เพิ่มข้อมูลสำเร็จแล้ว 
        </p>
      </Container>
    </div>
  );
}
export default Home;