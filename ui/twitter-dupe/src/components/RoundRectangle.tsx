import Paper from "@mui/material/Paper";

const RoundedRectangle = ({ children }: { children: React.ReactNode }) => {
  return (
    <Paper
      elevation={3}
      className="bg-gray-300 p-4 rounded-md shadow-md" // You can adjust the elevation to control the level of the shadow
    >
      {children}
    </Paper>
  );
};

export default RoundedRectangle;
