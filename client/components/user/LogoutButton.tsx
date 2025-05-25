'use client'
import { logout } from "@/app/login/actions"
import { useAuth } from "@/context/AuthContext"

export default function SubmitButton() {
  const { setUser } = useAuth()

  const handleLogout = () => {
    // clear user from context
    setUser(null)
    logout()
  }

  return (
    <button
      className="w-full bg-red-600 hover:bg-red-700 text-white font-medium py-2.5 rounded-lg transition-colors"
      type="submit"
      onClick={handleLogout}
    >
      Logout
    </button>
  )
}
