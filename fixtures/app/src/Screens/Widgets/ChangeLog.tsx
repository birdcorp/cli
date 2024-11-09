import React, { Suspense, lazy } from 'react'

import './ChangeLog.css'

const ReactMarkdown = lazy(() => import('react-markdown'))

const CHANGELOG = `
# Changelog

## [1.1.0] - 2024-05-18

### Added
- Introduced an optional \`color\` property to the \`LoadingIndicator\` component.
- The \`color\` property defaults to \`.white\` if no color is provided.
- Added examples in \`ContentView\` demonstrating the use of the \`LoadingIndicator\` with default and custom colors.

### Changed
- Modified the \`LoadingIndicator\` to use the provided \`color\` property or fall back to white if \`color\` is nil.

### Fixed
- Ensured consistent animation behavior regardless of the provided color.

## [1.0.0] - 2024-05-17

### Added
- Initial release of the \`LoadingIndicator\` component.
- Created a rotating circle animation using \`GeometryReader\` and \`ForEach\` to generate the circles.
- Implemented scaling and rotation effects to create a spinner animation.
`

export default function ChangeLog() {
  return (
    <div className="max-w-screen-sm mx-auto max-w-lg text-white p-3 brand-bg">
      <Suspense fallback={<div>Loading...</div>}>
        <div className="markdown-body">
          <ReactMarkdown>{CHANGELOG}</ReactMarkdown>
        </div>
      </Suspense>
    </div>
  )
}
