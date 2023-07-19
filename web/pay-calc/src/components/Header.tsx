import * as React from 'react';
import Avatar from '@mui/material/Avatar';
import Box from '@mui/material/Box';
import Grid from '@mui/material/Unstable_Grid2';
import { Link, Typography, Stack } from '@mui/material';
import AddHomeIcon from '@mui/icons-material/AddHome';
import AccountMenu from './Account';

export default function CenteredElementGrid() {

  const isAuthenticated: boolean = true

  return (
    <Box sx={{ flexGrow: 1 }}>
      <Grid container spacing={1} minHeight={100}>
        <Grid xs display="flex" justifyContent="left" alignItems="center">
          <AddHomeIcon sx={{color: "#f3cfd6"}} fontSize='large'/><Typography variant='h5'>Payment Calculator</Typography>
        </Grid>
        <Grid display="flex" justifyContent="center" alignItems="center">
        <Stack direction="row" spacing={2} fontSize={18}>
          <Link href='/'>Home</Link>
          <Link href='/records'>Records</Link>
          <Link href='/taxes'>Taxes</Link>
        </Stack>
        </Grid>
        <Grid xs display="flex" justifyContent="right" alignItems="center">
          {isAuthenticated && <AccountMenu/>}
        </Grid>
      </Grid>
    </Box>
  );
}