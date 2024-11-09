import { Link } from 'react-router-dom'

interface Props {
  to: string
  children: string | JSX.Element | JSX.Element[]
}

export function LinkMiniProgram({ to, children }: Props) {
  function onClick(e: any) {
    e.preventDefault()
    ;(window as any).miniProgram.call('navigateTo', {
      url: `${window.origin}${to}`,
    })
  }

  if (!!window?.bridge) {
    return (
      <button
        className="flex flex-row justify-between w-full"
        onClick={onClick}
      >
        {children}
      </button>
    )
  } else {
    return <Link to={to}>{children}</Link>
  }
}
