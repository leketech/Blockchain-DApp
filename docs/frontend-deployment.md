# Frontend Deployment Guide

This document explains how to build and deploy the dApp frontend for web, iOS, and Android platforms.

## Project Overview

The frontend is built with:
- React (TypeScript)
- Tailwind CSS for styling
- Material Symbols for icons
- Capacitor for mobile deployment

## Prerequisites

1. Node.js (v14 or higher)
2. npm or yarn
3. For iOS deployment: Xcode
4. For Android deployment: Android Studio

## Building for Web

### Development

To run the development server:

```bash
cd app
npm start
```

The app will be available at http://localhost:3000

### Production Build

To create a production build:

```bash
cd app
npm run build
```

The built files will be in the `build/` directory, which can be deployed to any web server.

## Building for Mobile Platforms

### iOS

To build for iOS:

```bash
cd app
npm run ios
```

This will:
1. Create a production build
2. Add the iOS platform to Capacitor
3. Open the project in Xcode

In Xcode, you can then:
1. Configure your signing certificate
2. Build and run on a simulator or device
3. Archive and submit to the App Store

### Android

To build for Android:

```bash
cd app
npm run android
```

This will:
1. Create a production build
2. Add the Android platform to Capacitor
3. Open the project in Android Studio

In Android Studio, you can then:
1. Build and run on an emulator or device
2. Generate a signed APK or App Bundle
3. Upload to the Google Play Store

## Project Structure

```
app/
├── public/
│   └── index.html          # Main HTML file
├── src/
│   ├── App.tsx             # Main app component
│   ├── index.tsx           # Entry point
│   ├── pages/
│   │   └── WelcomePage.tsx # Welcome screen
│   └── components/         # Reusable components
├── capacitor.config.ts     # Capacitor configuration
└── package.json            # Dependencies and scripts
```

## Customization

### Styling

The app uses Tailwind CSS with a custom configuration:

- Primary color: `#135bec`
- Light background: `#f6f6f8`
- Dark background: `#101622`
- Font: Space Grotesk

To modify the styling, you can edit the tailwind.config.js section in the HTML file.

### Content

To modify the content of the welcome screen, edit `src/pages/WelcomePage.tsx`.

### Icons

The app uses Material Symbols icons. To add new icons:

1. Add the icon name to the HTML: `<span class="material-symbols-outlined">icon_name</span>`
2. Ensure the Material Symbols font is loaded in the HTML head

## Deployment

### Web

The web version can be deployed to any static hosting service:
- AWS S3 + CloudFront
- Netlify
- Vercel
- GitHub Pages

### iOS App Store

1. Follow the iOS build process
2. In Xcode, go to Product > Archive
3. Follow Apple's guidelines for app submission

### Google Play Store

1. Follow the Android build process
2. In Android Studio, go to Build > Generate Signed Bundle / APK
3. Follow Google's guidelines for app submission

## Troubleshooting

### Common Issues

1. **Build errors**: Ensure all dependencies are installed with `npm install`
2. **iOS build failures**: Check Xcode version compatibility
3. **Android build failures**: Check Android Studio and SDK versions

### Getting Help

For issues with Capacitor, refer to the [Capacitor documentation](https://capacitorjs.com/docs).