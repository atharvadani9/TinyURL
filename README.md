# Tiny-URL

A full-stack URL shortening service built with React, TypeScript, Go, and PostgreSQL. This application allows users to create shortened URLs and automatically redirects users when they visit the shortened links.

![alt text](image.png)

## Features

- Create shortened URLs from long URLs
- Automatic redirection to original URLs
- Duplicate URL detection (returns existing short URL if original URL already exists)
- Clean, responsive Material-UI interface
- PostgreSQL database for persistent storage

## Technology Stack

### Frontend
- **React 19** with TypeScript
- **Material-UI (MUI)** for UI components
- **Vite** for build tooling and development server
- **Axios** for HTTP requests
- **React Router** for client-side routing

### Backend
- **Go 1.23** with Chi router
- **PostgreSQL** database
- **Goose** for database migrations
- **CORS** enabled for cross-origin requests


