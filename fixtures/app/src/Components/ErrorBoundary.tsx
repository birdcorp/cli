import React, { Component, ReactNode } from 'react'

interface ErrorBoundaryState {
  hasError: boolean
  errorMessage: string
  errorInfo: string
}

class ErrorBoundary extends Component<
  { children: ReactNode },
  ErrorBoundaryState
> {
  constructor(props: { children: ReactNode }) {
    super(props)
    this.state = { hasError: false, errorMessage: '', errorInfo: '' }
  }

  componentDidCatch(error: Error, info: React.ErrorInfo) {
    // Catch errors and update state
    this.setState({
      hasError: true,
      errorMessage: error.message ?? '',
      errorInfo: info.componentStack ?? '',
    })
  }

  resetError = () => {
    this.setState({ hasError: false, errorMessage: '', errorInfo: '' })
  }

  render() {
    const { hasError, errorMessage, errorInfo } = this.state

    if (hasError) {
      return (
        <div className="fixed inset-0 z-50 flex items-center justify-center bg-black bg-opacity-75 text-white p-4">
          <div className="bg-red-500 rounded-lg p-6 max-w-md text-left">
            <h2 className="text-xl font-bold mb-2">Something went wrong</h2>
            <p className="text-sm mb-2">
              <strong>Error:</strong> {errorMessage}
            </p>
            <pre className="text-xs whitespace-pre-wrap">{errorInfo}</pre>
            <button
              onClick={this.resetError}
              className="mt-4 px-4 py-2 bg-white text-black rounded hover:bg-gray-100"
            >
              Dismiss
            </button>
          </div>
        </div>
      )
    }

    return this.props.children
  }
}

export default ErrorBoundary
