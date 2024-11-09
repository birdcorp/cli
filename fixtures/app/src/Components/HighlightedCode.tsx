import { Prism as SyntaxHighlighter } from 'react-syntax-highlighter'
import { darcula } from 'react-syntax-highlighter/dist/esm/styles/prism'

interface Props {
  className?: string
  code: string
}

export const HighlightedCode: React.FC<Props> = ({ code, className = '' }) => {
  // Override background color of the darcula style
  const customStyle = {
    ...darcula,
    'pre[class*="language-"]': {
      ...darcula['pre[class*="language-"]'],
      background: '#202020', // Equivalent to bg-neutral-950
    },
  }

  return (
    <SyntaxHighlighter
      className={`${className} small-code`}
      language="javascript"
      style={customStyle} // Apply custom style here
    >
      {code}
    </SyntaxHighlighter>
  )
}
