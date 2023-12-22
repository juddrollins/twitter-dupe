import SignIn from "@/components/login";
import { Copyright } from "@mui/icons-material";
import {
  Box,
  Avatar,
  Typography,
  TextField,
  FormControlLabel,
  Checkbox,
  Button,
  Grid,
  Container,
} from "@mui/material";
import LockOutlinedIcon from "@mui/icons-material/LockOutlined";
import Link from "next/link";
import RoundedRectangle from "@/components/RoundRectangle";
import Post from "@/components/Post";

export default function Home() {
  return (
    <Container component="main" maxWidth="xs">
      <Box
        sx={{
          marginTop: 8,
          display: "flex",
          flexDirection: "column",
          alignItems: "center",
        }}
      >
        <Avatar sx={{ m: 1, bgcolor: "secondary.main" }}>
          <LockOutlinedIcon />
        </Avatar>
        <Typography component="h1" variant="h5">
          Home Page Holder
        </Typography>
        <Button variant="contained" href="/login">
          {" "}
          Login{" "}
        </Button>
      </Box>
    </Container>
  );
}
