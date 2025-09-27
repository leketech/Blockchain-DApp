# Blockchain dApp Web Interface

This is the web version of the Blockchain dApp interface.

## Features

- Responsive design for all screen sizes
- Dark/light mode support
- Wallet creation and import functionality
- Material Symbols icons
- Tailwind CSS styling

## Getting Started

### Prerequisites

- Node.js (v14 or higher)
- npm or yarn

### Installation

1. Install dependencies:
   ```bash
   npm install
   ```

2. Start the development server:
   ```bash
   npm start
   ```

3. Build for production:
   ```bash
   npm run build
   ```

## Project Structure

```
web/
├── public/
│   └── index.html          # Main HTML file
├── src/
│   ├── App.tsx             # Main app component
│   └── index.tsx           # Entry point
├── tailwind.config.js      # Tailwind CSS configuration
├── postcss.config.js       # PostCSS configuration
├── vite.config.ts          # Vite configuration
└── package.json            # Dependencies and scripts
```

## Customization

- Colors can be modified in the tailwind.config.js
- Fonts are loaded from Google Fonts (Space Grotesk)
- SVG icons can be customized in the component files

## Deployment

The app can be deployed to any static hosting service:
- AWS S3 + CloudFront
- Netlify
- Vercel
- GitHub Pages