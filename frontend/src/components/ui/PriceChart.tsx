import React from 'react';

interface ChartProps {
  dataPoints: number[];
  labels: string[];
}

export default function PriceChartPlaceholder({
  dataPoints,
  labels,
}: ChartProps) {
  // A modern, generative SVG chart placeholder

  // Calculate path
  const max = Math.max(...dataPoints);
  const min = Math.min(...dataPoints);
  const range = max - min;

  const width = 800;
  const height = 300;
  const padding = 20;

  const stepX = (width - padding * 2) / (dataPoints.length - 1);

  const points = dataPoints.map((val, idx) => {
    const x = padding + idx * stepX;
    // Invert Y axis for SVG
    const normalizedY = range === 0 ? 0.5 : (val - min) / range;
    const y = height - padding - normalizedY * (height - padding * 2);
    return `${x},${y}`;
  });

  const pathD = `M ${points.join(' L ')}`;

  // Gradient fill area
  const areaD = `${pathD} L ${width - padding},${height - padding} L ${padding},${height - padding} Z`;

  return (
    <div className="chart-wrapper">
      <div className="chart-header">
        <h3 className="chart-title">Market Value History</h3>
        <div className="chart-toggles">
          <button className="chart-toggle">1M</button>
          <button className="chart-toggle">3M</button>
          <button className="chart-toggle active">6M</button>
          <button className="chart-toggle">1Y</button>
          <button className="chart-toggle">ALL</button>
        </div>
      </div>

      <div className="svg-container">
        <svg
          viewBox={`0 0 ${width} ${height}`}
          className="price-chart"
          preserveAspectRatio="none"
        >
          <defs>
            <linearGradient id="chartGradient" x1="0" x2="0" y1="0" y2="1">
              <stop offset="0%" stopColor="#3b82f6" stopOpacity="0.4" />
              <stop offset="100%" stopColor="#3b82f6" stopOpacity="0" />
            </linearGradient>
            <linearGradient id="lineGradient" x1="0" x2="1" y1="0" y2="0">
              <stop offset="0%" stopColor="#3b82f6" />
              <stop offset="50%" stopColor="#8b5cf6" />
              <stop offset="100%" stopColor="#3b82f6" />
            </linearGradient>
          </defs>

          {/* Grid lines */}
          <line
            x1={padding}
            y1={padding}
            x2={width - padding}
            y2={padding}
            stroke="rgba(255,255,255,0.05)"
            strokeWidth="1"
          />
          <line
            x1={padding}
            y1={height / 2}
            x2={width - padding}
            y2={height / 2}
            stroke="rgba(255,255,255,0.05)"
            strokeWidth="1"
          />
          <line
            x1={padding}
            y1={height - padding}
            x2={width - padding}
            y2={height - padding}
            stroke="rgba(255,255,255,0.05)"
            strokeWidth="1"
          />

          {/* Area Fill */}
          <path d={areaD} fill="url(#chartGradient)" />

          {/* Main Line */}
          <path
            d={pathD}
            fill="none"
            stroke="url(#lineGradient)"
            strokeWidth="3"
            strokeLinecap="round"
            strokeLinejoin="round"
            className="chart-path"
          />

          {/* Data Points */}
          {points.map((pt, i) => {
            const [x, y] = pt.split(',');
            return (
              <circle
                key={i}
                cx={x}
                cy={y}
                r="4"
                fill="#0a0a0a"
                stroke="#8b5cf6"
                strokeWidth="2"
                className="chart-point"
              />
            );
          })}
        </svg>
      </div>

      <div className="chart-x-axis">
        {labels.map((label, i) => (
          <span key={i} className="axis-label">
            {label}
          </span>
        ))}
      </div>

      <style>{`
        .chart-wrapper {
          width: 100%;
          background: rgba(20, 20, 20, 0.4);
          border: 1px solid rgba(255, 255, 255, 0.08);
          border-radius: 16px;
          padding: 1.5rem;
        }
        
        .chart-header {
          display: flex;
          justify-content: space-between;
          align-items: center;
          margin-bottom: 1.5rem;
        }
        
        .chart-title {
          font-family: 'Outfit', sans-serif;
          font-size: 1.1rem;
          font-weight: 600;
          color: #ffffff;
        }
        
        .chart-toggles {
          display: flex;
          background: rgba(0,0,0,0.3);
          border-radius: 8px;
          padding: 0.25rem;
          border: 1px solid rgba(255,255,255,0.05);
        }
        
        .chart-toggle {
          background: transparent;
          border: none;
          color: #a1a1aa;
          padding: 0.25rem 0.75rem;
          border-radius: 4px;
          font-size: 0.75rem;
          font-weight: 600;
          cursor: pointer;
          transition: all 0.2s;
        }
        
        .chart-toggle:hover {
          color: #ffffff;
        }
        
        .chart-toggle.active {
          background: rgba(59, 130, 246, 0.2);
          color: #3b82f6;
        }
        
        .svg-container {
          width: 100%;
          height: 300px;
          position: relative;
        }
        
        .price-chart {
          width: 100%;
          height: 100%;
          overflow: visible;
        }
        
        .chart-path {
          stroke-dasharray: 2000;
          stroke-dashoffset: 2000;
          animation: drawPath 2s ease-out forwards;
        }
        
        @keyframes drawPath {
          to { stroke-dashoffset: 0; }
        }
        
        .chart-point {
          transition: all 0.2s;
          transform-origin: center;
        }
        
        .chart-point:hover {
          r: 6;
          fill: #8b5cf6;
          cursor: pointer;
        }
        
        .chart-x-axis {
          display: flex;
          justify-content: space-between;
          margin-top: 1rem;
          padding: 0 10px;
        }
        
        .axis-label {
          font-size: 0.75rem;
          color: #71717a;
        }
      `}</style>
    </div>
  );
}
