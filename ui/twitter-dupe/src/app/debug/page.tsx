// StickFigureRunning.js
'use client'
import React, { useState, useEffect } from 'react';
import './runner.css';

const StickFigureRunning = () => {
  const [duckPosition, setDuckPosition] = useState(-100);
  const [stickFigurePosition, setStickFigurePosition] = useState(0);

  useEffect(() => {
    const animationInterval = setInterval(() => {
      setDuckPosition((prev) => (prev < 100 ? prev + 1 : -100));
      setStickFigurePosition((prev) => (prev < 100 ? prev + 2 : 0));
    }, 50);

    return () => clearInterval(animationInterval);
  }, []);

  return (
    <div className="animation-container">
      <div className="duck" style={{ left: `${duckPosition}%` }}>
        🦆
      </div>
      <div className="stick-figure" style={{ left: `${stickFigurePosition}%` }}>
        <div className="head"></div>
        <div className="body"></div>
        <div className="arms">
          <div className="arm left-arm"></div>
          <div className="arm right-arm"></div>
        </div>
        <div className="legs">
          <div className="leg"></div>
          <div className="leg"></div>
        </div>
      </div>
    </div>
  );
};

export default StickFigureRunning;
