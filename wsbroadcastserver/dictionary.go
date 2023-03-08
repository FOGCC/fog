package wsbroadcastserver

import "compress/flate"

const DeflateCompressionLevel = flate.BestCompression

// The static dictionary was created by appending manually created dictionary to a dictionary generated with dictator tool (https://github.com/vkrasnov/dictator)
// * the dictator tool was used with default parameters, except for threshold which was set to 0.05% and compression level to 9
// * the input for the generator consisted of 61512 preprocessed messages gathered from public feed with wscat
// * the preprocessing substituted sequential numeric fields with pseudorandom values to minimize overfitting
func GetStaticCompressorDictionary() []byte {
	return []byte(`ADDUAtcQABTAABSAAEBCAoAEABUAABAOABAUACshAFYAACAKAAOAAHoSABAa/8AAAmJaAAAYAoAFABLAAAK/AASAEAAc3,"rACJA2,"ru4AgAoAgANCAAsAAATEtAOAAAQKgALXEAAB9AKAgAAAU6,"rABAe1,"rAWB4AEsAAAJAASDQ0,"rAVkAAFQAACQQA+AEQCAQAJAuA7u4AKABAGAQ5,"rACAIABAYADARAgAMAAAGgICA7,"rDu77A+gACACYloAFAiAFAcABzPUADgAEAW42AEABIBuAEAP/8AL68IAFAaABKABKgAgAwAQAaAAAAOuBAAIgAF9eEAGGoADu7gAEgAIAADgAgAAGAdzWUAKSxABkAGABAAGABAABAAAAEAAFAAEJoAAbANBuAFACR5AAAVgAGABwABAMAAAAKAYACgABYIAAIAQAwAADABgAQACgACADYAPAAA4AOjUpRAgAAIAqMA5U8BZoABIAFADgAAUAA7uAEeGjACQALAACADAAAH0AAgABAPj1gAAADAEABAPQkAKjAACcQ/AGAAoAHM9QArAAGACowAABhqAAQAD0JAcz1/pHKAAUAABwAgAKBvAAAgAC+vCABvwjrAAhpWEzBuA0AUAABoAgABASABwAAAAMNBECV6nsAACcQABAIABfXhAEAD4ACAoABAfQAAABYAASwAiAAQAKTG8zQSARgABOAYAEAoACBY0V4XYoAwAGgADAQAAMAYAC1xBAkACoA0AD4BIAHAgAJAG0AAgADAAYABYAHAKI3r14AkYTnKgAEcN5N+CAADQAANkZXYAABZIgxZFAAAADwAC7gALXmIPSAAAQCAoABiAAIABgASAD0JABASAOAQABaAIQAQAQgAACAFjRXhdigA9CQAABCaAABxr9SY0ACrSZH+AACBa8xB6AMAABgAWCIADEzCmjBgA//wAYABqAACGlYTAA5IqooZ8QADExvM0EABvBbWdOyAEAABAAMgyAbAANAwABa8xB6QADAAQABafqYnuuCGlYTNBAFAWAAojevXgACgAWAAPt8BfQ8rTGAFgAOQE5Fqv0},"signatureASADADMsADAEA1},"signatureAIACgAUxvoGihWXACoCkDEnMvAKI3r17/ARWORgkT0AAQAABAgAAMAUAADwAQACgAABACAMALQAEABoAAQSIQ6KAMywAMABoADnRyYW5zZmVydG8ueHl6AsaK8LsUxAOTwIQkpAZN8qtbsAEAAiAKo9n6OrkwrmNbAB0AxhKqWxTXUOAAGwsBnkLPj48OOEu4i+Bg7N6E+5VAHTHZNQbd9u7T+dx2rGTmwCxRolK/Aga49BeTw0N0p1oQEJKC3Yaf9tRACRhOcqAf5ASK/BwD55+H2iP6SaUDog581MAEjHetvV0nvbOaUOidaHT50hvTq4AAB3SUKdPZBHqWayikiVADJbrJ8bYAACAQAIWtLS6nz5xlyaIbTShd0hCeL/K4AAbABAAoAAQEAgADAQEAAMAADDB4RmVlRHluYW1pYwAACAAP+XCmGgSxyhSDSkP13kUz6921zADe3ewbDwKPaDF2XpISpo9ZkLoJ5AAQAH0CAnfW1TPgC5B2Eh5MaYaWckS0X9+wAElVuaXN3YXBWMwAD0FR7r+huch92SyCQ6GLG674wYEABoABBVdVUpiKOoBQS7rrExFnT8/UCtSwA89kS5/sYCs3qMaUnl9Ve4piKuQAQAIAEAEAAQAgq9JRH2KB+O9lb0NVvNSQVI/urEADAD7BUH1QAAnEADOXyS3qV6cun30tU6RG0o9yM2vb/AYAO7uAAEMgABQAIKvSUR9igfjvZW9DVbzUkFSP7qxD3OLoZC9eSRPkj6IynEJZwWwLEngAoAAAN4Lazp2QAIAIACQ5XAaZCZFijXCkSPVHiyhxfb0Rf/AAGABAq2r1ujOPoL1LZmYp/ZKkNKUqSpAAUABYABZede1RuOOQU9+mCJRS+RDpIAFKQANdPUlXVV5RM990ORf9SFSAALVdIAN4Lazp2Ab4KjwZ3A+NdFtyqiY9izeh/xiUYAAqApAxJzL4ADkiqihnAJGE5yoAG4AUAGwLajLDQl+uNV6F1uIx9i0eZdQYAKfOPF5skfWgUV1EKJj30I95rVlxAAAwAP+XCmGgSxyhSDSkP13kUz6921zIADO8ceRzdjD6pLWqzI5kRn9MOH/IUNXhV4AEj4AINa9v3iAWp0DXd+teAnNUAbhv/eF2tWGnv2BosQlRcnZADAAcCiawDB8N8AMAA5ACcUAAD6pxkA//ADAAO7uAkYTnKAXgAD/lwphoEscoUg0pD9d5FM+vdtcyADdlAGPVOVl2/yTn3xEoW4WP6qzMf/AAAD/lwphoEscoUg0pD9d5FM+vdtcyAAQAACAs0i4eyPVl34pSObzbKB+HslNcygACnOzOcP64nvt98ty2dAAsI/ImWCAFuJc82TOzDBVkBKDNbjzJ5ddBWAQAEgAKAAIKvSUR9igfjvZW9DVbzUkFSP7qx{"version":9,"confirmedSequenceNumberMessage":{"sAHdJQp09kEepZrKKSJUAMlusnxtADXT1JV1VeUTPfdDkX/UhUgAC1XSADekD4nEiiKHagpQt3fLCBSlWWsMAeUphNY1oRVlPlNwdsColK1tIFK3/ABvgqPBncD410W3KqJj2LN6H/GJRgAAUpPGyla4lBBAuNGPVX36gs9SAhYA2TtNC95pDBRPnyerIEHNRSXSCrQAMAQAABQBaWIT8MZSNWd8q7MyhQ96QDUnhoEgAcADZO00L3mkMFE+fJ6sgQc1FJdIKtAAIKvSUR9igfjvZW9DVbzUkFSP7qx/AG+EKg4f/8xJI9BLVZ9NKai/xs2gTrWw08RgdRwkAsXFzJEJyAJwUSAfdy+jvCYxYOoJuxbOGmuPwPqzaoAKWtxUhPmZf799QFuapip9iIg8NFAG5I8o79l7fnA6Dd8lKtG+LIm2nCsGj8N42qlSun8WPEoRYo9VpN9SOz76AOq0DN1iubhNjVSmCb9F2m7/oQYZ8},"signatureAABAQBkVBIWuv/+7I6lNbtx+8kngx0FlQAD//Ylj79H8alBkiEldlR1SY5iNJQAkIxNlNNJJHZfHtwiod0Jg5fFndT/{"version":3,"confirmedSequenceNumberMessage":{"s4},"signatureAEgAWAGUJhbGFuY2VyVjIAIAP+XCmGgSxyhSDSkP13kUz6921zIAGNSpWyq3E8eJc1sdZcPp2ijME5kAAt0wZ7sleWOhblY6MPXVyMUkzjz/AASoEqo8rwAgAAEAgAQABvhjoIxb4eXr7H1b0U9xQn0ehPPdAxTA97IpHlsgCsjHw7kloAQABEvMFpTAOqN/uGJin4KWfdSfwdhBtfkTCF2/AIgAAEAUDAMAAtcI7zczXsf6TOFOH+lOdEj8kPsALGivC7FAwAEAIBWANyK2RbSIeY12mcx41eEXTL5AHXBQ0YPbj4i9Dnf+UfiXJzLctLo/AAgAF/XEoC2OFFXspG5li8iFT/J55AojevXv/AMABABIAMAAAApLEAAAQADAAaACA5FAjtN8AR7e3r9jvvDqLS0HG9+I9xVD7G+/AABESUBLfAAADAVPCAkuMlYTGVTdV8BGR96LLnqakABU8ICS4yVhMZVN1XwEZH3osuepqAH9cdky8FPlmm4iDfKFJDMoXwxYHAQAQAC65B/tLcTkNxc0A5rgbfarONYGTAAAIKvSUR9igfjvZW9DVbzUkFSP7qxACKxyMEiegAACAgq9JRH2KB+O9lb0NVvNSQVI/urEAEgAwAACABAZQADUtgnWq4+DCQE2faPbO4IS1vrPd4QLkXpLZ8orAB5SmE1jWhFWU+U3B2wKiUrW0gUrQAAMAE29uZUluY2hWNUZlZUR5bmFtaWMAKAY7fVAJEWuQB32},"signature{"version":2,"confirmedSequenceNumberMessage":{"sA32luBRWVeSLpwGMgp9yI8aeEVv/ABgACwABYiv2fm5XR7gYZv4LhReKk8f4bjAEB9Aza7GVJX6XAVFxaQFIkIU41lPMNj/Aq0mR/yICgbE7KJ/+XCmGgSxyhSDSkP13kUz6921zIERERElTuslR3to+4Xtkp9zqWBYACbZCA4ePJOsM31XIxvp9CLoM935QA{"version":1,"confirmedSequenceNumberMessage":{"sAD//Ylj79H8alBkiEldlR1SY5iNJAbwW1nTs7gAAAKKFT7vZs+9m8X1HKE5/iZuVCTMADU8HwvHj7vCRbPJPvd7Hcxbd9KgJxABZede1RuOOQU9+mCJRS+RDpIAFKABLkor9fKd1x/Ts3ywAt+YIliq73P/EDk8IJdc8XejUTBxvBEPeWLNZ9oACOG8m/BAWSIMWRWAPvYSeYAf5vDzC1usVnfogjcZgJo/AOmB4qmknIKiL+XkTljIUjyNYieSAAAWAAFSMMLCvPObZbxpjr8BfD7ZklRef/ADEaB5sAjUXnsPU9PxN89FP3ZYBcA8xo+QZmHMVCAkOVwGmQmRYo1wpEj1R4socX29EX/AIALXCO83M17H+kzhTh/pTnRI/JD7AM5fJLepXpy6ffS1TpEbSj3Iza9vAABAAQAxGgebAI1F57D1PT8TfPRT92WAXA15NgUtHgltSMge85GPn9Y4QQhIALWXdrDcKM+Sq7Ql1unEWo/8pO7kACBRImhU8w7AhpWEzQAOVH+rTVzq/SniZTyxnmro7ZyFibC0nE5oAXTjMcsKf/OrWK/Jc41fiwAAAEAP//8ACIABAB0eRC3TSZRoGVjwxgoOOrkEg9CdwAcAIx2QGsLadAjqlu6XlqjEAB/kBIr8HAPnn4faI/pJpQOiDnzUwAEAIKvSUR9igfjvZW9DVbzUkFSP7qxABoCE/G94ZQABAfALNIuHsj1Zd+KUjm82ygfh7JTXMoC7gCYtAq+JB6qyr1N2unANhxiC/knUABIx3rb1dJ72zmlDonWh0+dIb06uAAAwABA5ZJCegrs6S3j7e4fGOAVfAWGFWQASMd629XSe9s5pQ6J1odPnSG9OrAFACQaEAMz7vou/DieyDam/2GfwcZEgaTn5Pj8hzxNi6xcVW3QCd95Wh9rpqDqFkJ/XHalZROe4ALfeM0QLcXEVmpcYy+dICGzs3ZaF/42ZSgzDPdb0gZ1HMiYld5Kv4jDsPGf+GgTkHuE+A81eBEZIe1JP3Eivas8mwHTazb37a1dLQsgUaAPl/TfdRF6eMGloNu4FK+SRYU5+0ARw3k34IAcRIbz3pk9L4p/F6Z8hyYJDRo7Z0AAgAP+XCmGgSxyhSDSkP13kUz6921zIAFgALyolQ7dqQWZUn3qrLnW+8K78Ww8AAIA4AWNFeF2K{"version":7,"confirmedSequenceNumberMessage":{"sAAUAcBqVcHoCkKyLkLNxno7lshA2CIMAAgALNIuHsj1Zd+KUjm82ygfh7JTXMABQAFgrEA5PAhCSAMMus2+Ib2OP/9g230TBJAdM/jWE/AABsC2oyw0JfrjVehdbiMfYtHmXUG/AEAC"signature":nulADAAABAQAABgAYAH+QEivwcA+efh9oj+kmlA6IOfNT/ACzSLh7I9WXfilI5vNsoH4eyU1zKAFTwgJLjJWExlU3VfARkfeiy56mpJzA14Q8QZJnvrOOF26BxNefMjl{"version":8,"confirmedSequenceNumberMessage":{"sACw1QLpOO1fTfLmgf5uQZ/yljHWK6AagACKXrr04N4ehYN0Nn3FQgUh2k0LjAAJxAMBgERERElTuslR3to+4Xtkp9zqWBYIAQAAAF0ImTqTM/Q8l0xC7igySsut8ILACD/lwphoEscoUg0pD9d5FM+vdtcyAKAA4ADHT+THFVEOwvjGHXDTl7MgQ/VavADAAEAAQEAAIT8b3hlAtf47wu9MNMwjOSLf8vyxsb+Jo47/AESVPtsRLrAvtKFTnb5BkMJfAFkiDFkVAGizRlgz+3KnDs30heDkx72GZfxFAACE/G94ZQAWACCr0lEfYoH472VvQ1W81JBUj+6sQAIOYloCUhYWhD1n9Tdbn1eGSVNWnkdwl8/SAuIQhehPpAABV9KG/xlXPVe0yXyM4od7oT3VN8A9q7TL8R03VcXEF26XqVyaPRusANEMgz9DBeEFOmS8c4xVA4H0gQTKAIAAEAHSAIk7CuY8ozuN0DICUCi9Im/NSAABihBVWWwAL+i9ov5rWDcPPsc7wRzDrf6JRQkACZDrKON4ZZuTop1G/0HwjcYxbdmAAAOQE5FqvAHgBJeD8AigCFGT4AAD0AzVQAPHEAAmmZkQIQQihvCOpzNUNbWPlMJjbAAkAAAmACajPrPUT+z1eOfWVLIYI6YWz3G7ADAAEBAQA7gAP+XCmGgSxyhSDSkP13kUz6921zAJrFJ5AT7f7HTFwpdvyDGtBSdALAlQFuNq204BUXNc7TmSp/pU4WvQADNrsZUlfpcBUXFpAUiQhTjWU8w2P/C28epj6Vm//ofq9WpKx7n937ar8AIcM+N1dnI6xQD39bmpHU/TWF6UADeC2s6dkAJjp0oh0ODnpaoAFprUcdwu/d4jAGAAQAx0rKuMCjQPWF0AjLUh1k0lVBcaACVStx0SBY0tNJ4xFmFO05swXro0ADPEt0011l9Bv+Y+F0rlIPZ1ffZUAAP+XCmGgSxyhSDSkP13kUz6921zIAKJuhcBwvCIzw+vizcFJGMxy4qDkACAOA9zi6GQvXkkT5I+iMpxCWcFsCxJ4{"version":0,"confirmedSequenceNumberMessage":{"sAD9CGvHzVxIHcychevkeKHAtp/LuQAAJeZtZ4TPZWo8UW+NF8l6O6XgVzgADpgeKppJyCoi/l5E5YyFI8jWInkAI4byb8EA9uk+sohljeXi6YL5nSs3iyKVnRUAPWtfzeC6KZ7/6KXaE4nz5/MeBvhC7ACCR+876cyqwM2g9E2EO/fZbiHL0A0CalPeZU0OuylKwRI/1NyKC1Bf0AIAAQABAEAD8WhpusHaix60G7SLJDX5xDjWtCgAcACQABkiw287x2LKMAtKRrshAvhLFoSrAAFbU9J3r4YPUcPmhD+AdQBwJruzoijYukRZWKdaBwTVZr8siAuQHEUru+KQAAgAgAFAgq9JRH2KB+O9lb0NVvNSQVI/urEAGwLajLDQl+uNV6F1uIx9i0eZdQb/A=="},"delayedMADAAEAAQAAMFg5P6Wf4rae1jZxXgmfrvsHnpjAOu6Rn7LayEjkXgDMYnOrifKEurfAABuAAAooVPu9mz72bxfUcoTn+Jm5UJMABTm94Nfb0za3kUiqdCiDGYu/YDQgAAEAICAgq9JRH2KB+O9lb0NVvNSQVI/urEAH/xzLaL26dBZQ9VIYtTCYs0qTPl/ADGAC+wACzQACswAAGkBBgq9JRH2KB+O9lb0NVvNSQVI/urHQ4w2wAKCGCjLsAKAGRUEha6//7sjqU1u3H7ySeDHQWVBTm94Nfb0za3kUiqdCiDGYu/YDQ"timestampAAJDlcBpkJkWKNcKRI9UeLKHF9vRF/A3gtrOnZAbCwGeQs+Pjw44S7iL4GDs3oT7lUAQACAACuOSh5oiiySE2bH4Cl0LcID+ecIAAPrjnsCXMMoPFCYqY20tfFU5NTdSBAAVgAKtJkf4AgAIA2DG6+AiA5HHnAWAD4nXuchk9Ym79TqCEFEHYis16qQAAjhvJvwQAzLAiATBPSHaQdLx1jNLgPMkHKctK3ADYMbr4AERERElT7bES6wL7ShU52+QZDCX0APl8cHAk7w3T53oIJFVaRrYiv7UA/6},"signatureAC3TBnuyV5Y6FuVjow9dXIxSTOPP/AA/Qhrx81cSB3MnIXr5HihwLafy7kADAAABAAEAABWQAFZAABAFAAwACqVK09MK93tg2TmuNW5mBt6aTaZ1g/Arli0tPy5IFITp077175656e6AhPxveGUAPQkAAFpYhPwxlI1Z3yrszKFD3pANSeGjuAIaVhM0ABQAAKJrAMHw3wAwADkAJxQAAPqnGQAKAifyGhNNeCrXtzPa079pf6i+N4tArgAGAMpjCZphW4lCePsTuZxCIFhPiegFAAQxERYdwuskWg9Ru3kxDB6A0BKbQABvFe6SWKzevzVtt6tge7JVoAxv3AU5veDX29M2t5FIqnQogxmLv2A0IACAAwAxniIt5GKslZuvKuyEhpeuwrvXcAADlkkJ6CuzpLePt7h8Y4BV8BYYVZAQAFgACAABAZgyCUcqe5hdiHOTTyjyV4aV/9n4AU5veDX29M2t5FIqnQogxmLv2A0ADAAABAQEAADIMAKgKQMScy+Aca/UmNASVW5pc3dhcFYzAnBRIB93L6O8JjFg6gm7Fs4aa4/A+rNq/5cKYaBLHKFINKQ/XeRTPr3bXMgA5PAhCSkAUACCr0lEfYoH472VvQ1W81JBUj+6sQABAAFgAM5fJLepXpy6ffS1TpEbSj3Iza9v/7},"signatureA5ZJCegrs6S3j7e4fGOAVfAWGFWAoAfl4WjWiQXe4U7d71xZZYDiJ6y08AKAAOJjNTmGpGOBRMQeRM66ydCnbsqzAADA/5cKYaBLHKFINKQ/XeRTPr3bXMADGeIi3kYqyVm68q7ISGl67Cu9dwA5},"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:{"kind":9,"senARElAS3wACAyuajp9+dnaMqScQA9PkPhYWsf/ACC1wjvNzNex/pM4U4f6U50SPyQ+yAALXCO83M17H+kzhTh/pTnRI/JD7AIAGkrL60MAPkJFvh6lLacjLzZCjnY09PM8KPpaFEcOFag+EKg4f/8xJI9BLVZ9NKai/xs2gTrWw08RgdRwkAsXFzJEJyABos0ZYM/typw7N9IXg5Me9hmX8Rf/Ag3zFstBFvjZlE1zrbcrw902iWbt0PCyFIQBoXji4aCwxjL1Xx4bOyw9gvQe5HFrtMAPD12Eza/BQVgbuHV6T7iAtJxOaAF04zHLCn/zq1ivyXONX4uAgASS4WBg/AQAgq9JRH2KB+O9lb0NVvNSQVI/urEABgAAKADALyolQ7dqQWZUn3qrLnW+8K78Ww8572","blockA47tanGk3anX93aUbjz6TsQ3ga5v/AJ3abvPZGcm8iIXVVgmZo2QEMejm/AMAAE29uZUluY2hWNUZlZUR5bmFtaWAMbZc7MbsTXKuoPPBXTANHvXY+zFAGABgADAE:"0xa4b1e63cb4901e327597bc35d36fe8a23e4c253f","blockAGRW5ZaIAORx5wFoAIAADAFiK/Z+bldHuBhm/guFF4qTx/huM6e672","blockCuCr0lEfYoH472VvQ1W81JBUj+6sQAB9P+XCmGgSxyhSDSkP13kUz6921zICOG8m/BAIEkuFgYPwAbwAJ3zyoBAPuIp8ONbW1K3DbBcRulL3TJXjMGAAZS0nwPcncc5cdv1ADt1htAasbZf/AF0ImTqTM/Q8l0xC7igySsut8ILAwWDk/pZ/itp7WNnFeCZ+u+weemMABvAAEAwC8qJUO3akFmVJ96qy51vvCu/FsPAIAACijYukRZWKdaBwTVZr8siAuQLkuVysKPtebQwd/tK6AA+8BAq43zYVrDKcAAGQ0NZSrm1bpkIe/pvIzXbJMLR8XAKoAAHAAJy9XQfdDOzGYWH8k9fJAA2hALoSIiIiIo2LpEWVinWgcE1Wa/LI/ALXCO83M17H+kzhTh/pTnRI/JD7AFOb3g19vTNreRSKp0KIMZi79gNCAACihU+72bPvZvF9RyhOf4mblQkzsUEBgq9JRH2KB+O9lb0NVvNSQVI/urEABC4afU0IOYloCUq16ZB83Bz9DMa5ve9rFZcqRFuD+AuIQ6XcAIAABos0ZYM/typw7N9IXg5Me9hmX8RQAAqj2fo6uTCuY1sAHQDGEqpbFNdQ4AAnBRIB93L6O8JjFg6gm7Fs4aa4/A+rNqAQAAEAAD/lwphoEscoUg0pD9d5FM+vdtcyADjX6kxoANgxuvABa8xB6QAA//2JY+/R/GpQZIhJXZUdUmOYjSUAB1VWkVlZjzcCvdff9iM6MXwVbT3f/AERJQEt8:{"kind":13,"senAuhIiA2hAAnL1dB90M7MZhYfyT18kADaEAAgq9JRH2KB+O9lb0NVvNSQVI/urEAG8V7pJYrN6/NW23q2B7slWgDG/fARElAS3AiATBPSHaQdLx1jNLgPMkHKctK3AB+AyWylIdckeoGKCbCxXDjj5Yoo/AUAcBqVcHoCkKyLkLNxno7lshA2CIMijYukRZWKdaBwTVZr8siAuQMEuVysKGRUEha6//7sjqU1u3H7ySeDHQWVAAEAAHVVaRWVmPNwK919/2IzoxfBVtPd/ABbZ3bNnFF2j8kVyten6PXEpjMRMf/ACQjE2U00kkdl8e3CKh3QmDl8Wd1P/B7AiMAkaftASMAcvcAagBNYKjU5x1Zm4EEJQ8ABdCJk6kzP0PJdMQu4oMkrLrfCCwA2hAAnL1dB90M7MZhYfyT18kADaAaLNGWDP7cqcOzfSF4OTHvYZl/EX/ABsLAZ5Cz4+PDjhLuIvgYOzehPuVQAAAgq9JRH2KB+O9lb0NVvNSQVI/urAdVVpFZWY83Ar3X3/YjOjF8FW093/Af5ASK/BwD55+H2iP6SaUDog581P/AM2uxlSV+lwFRcWkBSJCFONZTzDY/AYjg3OdaN0PhEED2437Baft7Vu+YAAgACABACAM7xx5HN2MPqktarMjmRGf0w4f8hQ1eFXgABlLSfA9ydxzlx2/UAO3WG0Bqxtl/AONfqTGgA3vHA3tm+x/GhZwgZgzJA8CeyXv8AASTvXas6AW4lzzZM7MMFWQEoM1uPMnl10FYAFgAORQI7TfABAWACBJLhYGD8AAC8qJUO3akFmVJ96qy51vvCu/FsPA41+pMaAALX+O8LvTDTMIzki3/L8sbG/iaOO/APjZlKDMM91vSBnUcyJiV3kq/iMOw8Z/4aBOQe4T4DzV4ERkh7Uk/cSK9qzybAdNrNvftrV0tCyBRrigAEk712rOgAN7xwN7ZvsfxoWcIGYMyQPAnsl7/AAMAvKiVDt2pBZlSfeqsudb7wrvxbDwAIAAgAEAYAAEEiEOigsRead":1AD5fHBwJO8N0+d6CCRVWka2Ir+1AP/AQAA+Qi6+QE6lN7MDAnDtfbpLvQYQSXVZIpm41KY4aA0Zg/IrzBEZFKfSKd44D0D5NNLzV+bbwz7880jjGQvf7kBAh7uALZwOND8AUQAAcpAxzgC/JABIACAGUtJ8D3J3HOXHb9QA7dYbQGrG2X/AJCMTZTTSSR2Xx7cIqHdCYOXxZ3U/Af/HMtovbp0FlD1Uhi1MJizSpM+UAAMBZede1RuOOQU9+mCJRS+RDpIAFKQAdHkQt00mUaBlY8MYKDjq5BIPQncA0BaG9RrhcXAKUg/vs8vPmEQJuRQAAQABALcsF3V6OPkC2pAwAFl517VG445BT36YIlFL5EOkgAUpADAWXnXtUbjjkFPfpgiUUvkQ6SABSkADlj/8nf9EPCqU8hsSnUKYkeMuwY/ABESUBLfAEAIA+XxwcCTvDdPneggkVVpGtiK/tQD/ABJO9dqzoAd2enhQvP4TZCvGp8sthfH4IQgpAB/kBIr8HAPnn4faI/pJpQOiDnzU/AIAD/lwphoEscoUg0pD9d5FM+vdtcyAAD4eZRwGpVwegKQrIuQs3GejuWyEDYIg+GgjT7g32pLfoKn8gp2PxxoJuYXYyPmVa9k8yMYgn0hEtS4AFO/gzpdbE3aiI9pwiyIyfNWpBYU/AWvMQekAAJiK5REozb4+ZRNc623K8PdNolm7dDwshSEAaF44uGgsMYy9V8eGzssPYL0HuRxa7TADw9dhM2vwUFYG7h1ek+4wAC8qJUO3akFmVJ96qy51vvCu/FsPAewIjAJGn7QEjAHL3AGoATWCo1OcdWZuBBCUPABBIhDooAD4m5R/XHZMvBT5ZpuIg3yhSQzKF8MWB/hjoN3yUq0b4sibacKwaPw3jaqVK6fxY8ShFij1Wk31I7PvAHsCIwCRp+0BIwBy9wBqAE1gqNTnHVmbgQQlDwAomsAwfDfADAAOQAnFAAA+qcZADaEACcvV0H3QzsxmFh/JPXyQANoQA5FAjtN8ACBrj0F5PDQ3SnWhAQkoLdhp/21HACAAIABAGAAORQI7TfADkBORarwAXQiZOpMz9DyXTELuKDJKy63wgsAJDlcBpkJkWKNcKRI9UeLKHF9vRF/AIGuPQXk8NDdKdaEBCSgt2Gn/bUcAAAABAALyolQ7dqQWZUn3qrLnW+8K78Ww8ijYukRZWKdaBwTVZr8sj/AIgEwT0h2kHS8dYzS4DzJBynLStwAD9CGvHzVxIHcychevkeKHAtp/LuQg6b3tXySkFo5PO/ROACmMmTsiN2qtjFjH3alxilTL6oK5AqAAbxXuklis3r81bberYHuyVaAMb98AvKiVDt2pBZlSfeqsudb7wrvxbDwAO7uACIBME9IdpB0vHWM0uA8yQcpy0rcACAAQAD7Xm0MHf7SugAPvAQKuN82FawynADu7AEEiEOiAEAAgAD9CGvHzVxIHcychevkeKHAtp/LusRead":A+15tDB3+0roAD7wECrjfNhWsMpwAPtebQwd/tK6AA+8BAq43zYVrDKcAIADkBORarAG8FtZ07ADO8ceRzdjD6pLWqzI5kRn9MOH/IUNXhV4aBpOfk+PyHPE2LrFxVbdAJ33laH2umoOoWQn9cdqVlE57iAFrzEHpAANoQAJy9XQfdDOzGYWH8k9fJAA2hABAAEAB4ASXg/AIoAhRk+AAA9AM1UADxxAAC6EiIiIiKNi6RFlYp1oHBNVmvyyP/A7uu4AAC0rqqxAlj0AAGizRlgz+3KnDs30heDkx72GZfxF/AD8WhpusHaix60G7SLJDX5xDjWtCABAAIAEAAQ:{"kind":12,"senABsC2oyw0JfrjVehdbiMfYtHmXUG/A/FoabrB2osetBu0iyQ1+cQ41rQoADlR/q01c6v0p4mU8sZ5q6O2chYmwAQAABA/Qhrx81cSB3MnIXr5HihwLafy7kABSFwmzzX8H4pcivguiiozg6AbbwwAza7GVJX6XAVFxaQFIkIU41lPMNgAQAACAEABYAbwADUtgnWq4+DCQE2faPbO4IS1vrPdAQSIQ6KAPxaGm6wdqLHrQbtIskNfnEONa0KAgAOQE5FqvACA5ATkWq8AGACgADkUCO03wAP0Ia8fNXEgdzJyF6+R4ocC2n8u5ACAAmAFIXCbPNfwfilyK+C6KKjODoBtvDAgFvJ9vLDqfIXmm48gAMU7XBXIGAKZDYfoI9PMqDaDIH3LLVZg3r83YAD/lwphoEscoUg0pD9d5FM+vdtcyAAUhcJs81/B+KXIr4LooqM4OgG28MAzvHHkc3Yw+qS1qsyOZEZ/TDh/yFDV4VeAAIKvSUR9igfjvZW9DVbzUkFSP7qxAA/5cKYaBLHKFINKQ/XeRTPr3bXMgACCr0lEfYoH472VvQ1W81JBUj+6sQACCr0lEfYoH472VvQ1W81JBUj+6sA/5cKYaBLHKFINKQ/XeRTPr3bXM"restId":null,"baseFeeL1AP+XCmGgSxyhSDSkP13kUz6921zIAgq9JRH2KB+O9lb0NVvNSQVI/urEAD/lwphoEscoUg0pD9d5FM+vdtcysReadAIKvSUR9igfjvZW9DVbzUkFSP7qxl},"l2Msg":"AwAe":{"head:"0xa4b072","block:{"kind":3,"sen073657175656e6estId":null,"baseFeeL10,"message":{"message":{"header":{"kind":13,"sender":"0xa4b
1,"message":{"message":{"header":{"kind":13,"sender":"0xa4b 2,"message":{"message":{"header":{"kind":13,"sender":"0xa4b 3,"message":{"message":{"header":{"kind":13,"sender":"0xa4b 4,"message":{"message":{"header":{"kind":13,"sender":"0xa4b 5,"message":{"message":{"header":{"kind":13,"sender":"0xa4b 6,"message":{"message":{"header":{"kind":13,"sender":"0xa4b 7,"message":{"message":{"header":{"kind":13,"sender":"0xa4b 8,"message":{"message":{"header":{"kind":13,"sender":"0xa4b 9,"message":{"message":{"header":{"kind":13,"sender":"0xa4b 0,"message":{"message":{"header":{"kind":12,"sender":"0xa4b 1,"message":{"message":{"header":{"kind":12,"sender":"0xa4b 2,"message":{"message":{"header":{"kind":12,"sender":"0xa4b 3,"message":{"message":{"header":{"kind":12,"sender":"0xa4b 4,"message":{"message":{"header":{"kind":12,"sender":"0xa4b 5,"message":{"message":{"header":{"kind":12,"sender":"0xa4b 6,"message":{"message":{"header":{"kind":12,"sender":"0xa4b 7,"message":{"message":{"header":{"kind":12,"sender":"0xa4b 8,"message":{"message":{"header":{"kind":12,"sender":"0xa4b 9,"message":{"message":{"header":{"kind":12,"sender":"0xa4b 0,"message":{"message":{"header":{"kind":9,"sender":"0xa4b 1,"message":{"message":{"header":{"kind":9,"sender":"0xa4b 2,"message":{"message":{"header":{"kind":9,"sender":"0xa4b 9,"message":{"message":{"header":{"kind":9,"sender":"0xa4b 4,"message":{"message":{"header":{"kind":9,"sender":"0xa4b 5,"message":{"message":{"header":{"kind":9,"sender":"0xa4b 6,"message":{"message":{"header":{"kind":9,"sender":"0xa4b 7,"message":{"message":{"header":{"kind":9,"sender":"0xa4b 8,"message":{"message":{"header":{"kind":9,"sender":"0xa4b 9,"message":{"message":{"header":{"kind":9,"sender":"0xa4b 0,"message":{"message":{"header":{"kind":3,"sender":"0xa4b 1,"message":{"message":{"header":{"kind":3,"sender":"0xa4b 2,"message":{"message":{"header":{"kind":3,"sender":"0xa4b 3,"message":{"message":{"header":{"kind":3,"sender":"0xa4b 4,"message":{"message":{"header":{"kind":3,"sender":"0xa4b 5,"message":{"message":{"header":{"kind":3,"sender":"0xa4b 6,"message":{"message":{"header":{"kind":3,"sender":"0xa4b 7,"message":{"message":{"header":{"kind":3,"sender":"0xa4b 8,"message":{"message":{"header":{"kind":3,"sender":"0xa4b 9,"message":{"message":{"header":{"kind":3,"sender":"0xa4b 1,"timestamp": 2,"timestamp": 3,"timestamp": 4,"timestamp": 5,"timestamp": 6,"timestamp": 7,"timestamp": 8,"timestamp": 9,"timestamp": 0,"requestId":null,"baseFeeL1":null},"l2Msg":" 1,"requestId":null,"baseFeeL1":null},"l2Msg":" 2,"requestId":null,"baseFeeL1":null},"l2Msg":" 3,"requestId":null,"baseFeeL1":null},"l2Msg":" 4,"requestId":null,"baseFeeL1":null},"l2Msg":" 5,"requestId":null,"baseFeeL1":null},"l2Msg":" 6,"requestId":null,"baseFeeL1":null},"l2Msg":" 7,"requestId":null,"baseFeeL1":null},"l2Msg":" 8,"requestId":null,"baseFeeL1":null},"l2Msg":" 9,"requestId":null,"baseFeeL1":null},"l2Msg":" 1,"requestId":"0x00000000000000000000000000000000000000000000000000000000000 2,"requestId":"0x00000000000000000000000000000000000000000000000000000000000 3,"requestId":"0x00000000000000000000000000000000000000000000000000000000000 4,"requestId":"0x00000000000000000000000000000000000000000000000000000000000 5,"requestId":"0x00000000000000000000000000000000000000000000000000000000000 6,"requestId":"0x00000000000000000000000000000000000000000000000000000000000 7,"requestId":"0x00000000000000000000000000000000000000000000000000000000000 8,"requestId":"0x00000000000000000000000000000000000000000000000000000000000 9,"requestId":"0x00000000000000000000000000000000000000000000000000000000000 {"version":1,"confirmedSequenceNumberMessage":{"sequenceNumber":
{"version":1,"messages":[{"sequenceNumber":
","blockNumber":
0,"timestamp":
0,"requestId":"0x00000000000000000000000000000000000000000000000000000000000
","baseFeeL1":null},"l2Msg":"
=="},"delayedMessagesRead":
},"signature":null}]}
`)
}
