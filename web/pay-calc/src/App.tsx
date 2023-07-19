import React from 'react';
import { BrowserRouter, Routes, Route } from 'react-router-dom';
import MainPage from './pages/MainPage';
import RecordPage from './pages/RecordPage';
import TaxesPage from './pages/TaxesPage';
import SignUpPage from './pages/SignUpPage';
import { Container } from '@mui/material';


function App() {
  return (
    <Container maxWidth="lg" sx={{bgcolor: '#efefef', borderRadius: 4, p: 1}}>
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<MainPage />} />
        <Route path="/records" element={<RecordPage />} />
        <Route path="/taxes" element={<TaxesPage />} />
        <Route path="/signup" element={<SignUpPage />} />
      </Routes>
    </BrowserRouter>
    </Container>
  );
}

export default App;
