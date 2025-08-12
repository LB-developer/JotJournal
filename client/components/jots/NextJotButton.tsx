"use client";

interface Props {
    monthNumberAsString: string;
    forward?: boolean;
}

export default function NewComponent({ monthNumberAsString, forward }: Props) {
    const handleMoveToNextJot = async () => {
        let monthNumber = Number(monthNumberAsString);
        if (forward == false) {
            monthNumber -= 1;
        } else {
            monthNumber += 1;
        }

        const monthStringWithPadded0 = monthNumber.toString().padStart(2, "0");
        const url = "/dashboard?month=" + monthStringWithPadded0;
        window.location.href = url;
    };
    return (
        <>
            <button className="hover:bg-blue-200" onClick={handleMoveToNextJot}>
                {forward ? ">" : "<"}
            </button>
        </>
    );
}
