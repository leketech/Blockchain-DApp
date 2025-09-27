# Blockchain dApp Frontend

This is the frontend interface for the Blockchain dApp, built with React and TypeScript.

## Features

- Welcome screen with wallet creation and import options
- Dark/light mode support
- Responsive design for web, iOS, and Android
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
src/
  ├── App.tsx          # Main app component
  ├── index.tsx        # Entry point
  ├── pages/           # Page components
  │   └── WelcomePage.tsx  # Welcome screen
  └── components/      # Reusable components
```

## Mobile Integration

To integrate with iOS and Android platforms:

1. For iOS: Use Capacitor or React Native to wrap the web app
2. For Android: Use Capacitor or React Native to wrap the web app

## Customization

- Colors can be modified in the tailwind.config.js
- Fonts are loaded from Google Fonts (Space Grotesk)
- SVG icons can be customized in the component files

## Deployment

The app can be deployed to:
- Web servers
- iOS App Store (via Capacitor/Cordova)
- Google Play Store (via Capacitor/Cordova)