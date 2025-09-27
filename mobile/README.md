# Blockchain dApp Mobile

This is the mobile version of the Blockchain dApp, built with React Native.

## Features

- Wallet creation and import functionality
- Cross-platform support (iOS and Android)
- Native mobile UI components
- Responsive design for all device sizes

## Prerequisites

- Node.js (v14 or higher)
- npm or yarn
- Expo CLI
- For iOS: Xcode
- For Android: Android Studio

## Installation

1. Install dependencies:
   ```bash
   npm install
   ```

2. Start the development server:
   ```bash
   npm start
   ```

3. To run on iOS:
   ```bash
   npm run ios
   ```

4. To run on Android:
   ```bash
   npm run android
   ```

## Project Structure

```
mobile/
├── src/
│   ├── App.js              # Main app component
│   ├── components/         # Reusable components
│   ├── screens/            # Screen components
│   └── services/           # Service utilities
├── app.json                # Expo configuration
├── babel.config.js         # Babel configuration
└── package.json            # Dependencies and scripts
```

## Building for Production

### iOS

To build for iOS:
```bash
expo build:ios
```

### Android

To build for Android:
```bash
expo build:android
```

## Customization

- Colors can be modified in the StyleSheet in App.js
- UI components can be customized in the components directory
- Navigation can be added in the screens directory

## Deployment

### iOS App Store

1. Build the app:
   ```bash
   expo build:ios
   ```

2. Follow Apple's guidelines for app submission

### Google Play Store

1. Build the app:
   ```bash
   expo build:android
   ```

2. Follow Google's guidelines for app submission