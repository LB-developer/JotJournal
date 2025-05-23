import { logout } from "@/app/login/actions"

export default async function SubmitButton() {
    
    return (
        <button 
          className="w-full bg-red-600 hover:bg-red-700 text-white font-medium py-2.5 rounded-lg transition-colors"
          type="submit"
          onClick={logout}
        >
          Logout
        </button>
    )
}
